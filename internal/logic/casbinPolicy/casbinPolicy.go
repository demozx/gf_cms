package casbinPolicy

import (
	"database/sql"
	"fmt"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
	"sync"
	"time"

	sqlAdapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	insPolicy            = sCasbinPolicy{}
	ObjBackend    string = "backend"
	ObjBackendApi string = "backend_api"

	casbinOnce      sync.Once
	casbinEnforcer  *casbin.Enforcer
	casbinInitError error
)

type sCasbinPolicy struct{}

func init() {
	service.RegisterCasbinPolicy(New())
}

func New() *sCasbinPolicy {
	return &sCasbinPolicy{}
}

func CasbinPolicy() *sCasbinPolicy {
	return &insPolicy
}

// 初始化casbin
func initCasbin() *casbin.Enforcer {
	casbinOnce.Do(func() {
		casbinEnforcer, casbinInitError = newCasbinEnforcer()
	})
	if casbinInitError != nil {
		panic(casbinInitError)
	}
	return casbinEnforcer
}

func newCasbinEnforcer() (*casbin.Enforcer, error) {
	var dbType = util.Util().GetConfig("database.default.type")
	var dbUser = util.Util().GetConfig("database.default.user")
	var dbPass = util.Util().GetConfig("database.default.pass")
	var dbHost = util.Util().GetConfig("database.default.host")
	var dbPort = util.Util().GetConfig("database.default.port")
	var dbName = util.Util().GetConfig("database.default.name")
	var dbPrefix = util.Util().GetConfig("database.default.prefix")

	db, err := sql.Open(dbType, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	db.SetMaxOpenConns(gconv.Int(util.Util().GetConfig("database.default.maxOpen")))
	db.SetMaxIdleConns(gconv.Int(util.Util().GetConfig("database.default.maxIdle")))
	db.SetConnMaxLifetime(time.Second * 10)

	a, err := sqlAdapter.NewAdapter(db, dbType, dbPrefix+"rule_permissions")
	if err != nil {
		return nil, fmt.Errorf("create casbin sql adapter: %w", err)
	}
	if err = sanitizeCasbinRuleTable(db, dbPrefix+"rule_permissions"); err != nil {
		return nil, fmt.Errorf("sanitize casbin rule table: %w", err)
	}

	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, fmt.Errorf("new enforcer: %w", err)
	}

	// Load the policy from DB.
	if err = e.LoadPolicy(); err != nil {
		log.Println("LoadPolicy failed, err: ", err)
	}

	return e, nil
}

func sanitizeCasbinRuleTable(db *sql.DB, tableName string) error {
	query := fmt.Sprintf(
		"UPDATE %s SET p_type=COALESCE(p_type,''), v0=COALESCE(v0,''), v1=COALESCE(v1,''), v2=COALESCE(v2,''), v3=COALESCE(v3,''), v4=COALESCE(v4,''), v5=COALESCE(v5,'') WHERE p_type IS NULL OR v0 IS NULL OR v1 IS NULL OR v2 IS NULL OR v3 IS NULL OR v4 IS NULL OR v5 IS NULL",
		tableName,
	)
	_, err := db.Exec(query)
	return err
}

// ObjBackend 获取后台obj
func (*sCasbinPolicy) ObjBackend() string {
	return ObjBackend
}

// ObjBackendApi 获取后台接口obj
func (*sCasbinPolicy) ObjBackendApi() string {
	return ObjBackendApi
}

// CheckByRoleId 检测角色权限
func (*sCasbinPolicy) CheckByRoleId(roleId, obj, act string) bool {
	has, err := initCasbin().Enforce(roleId, obj, act)
	if err != nil {
		log.Println("Enforce failed, err: ", err)
		g.Log().Line(true).Async(true).Error(util.Ctx, "Enforce failed, err: ", err)
	}
	if !has {
		g.Log().Line(true).Async(true).Warning(util.Ctx, "没有操作权限："+roleId+","+obj+","+act)
	} else {
		return true
	}
	return false
}

// CheckByAccountId 检测用户权限
func (*sCasbinPolicy) CheckByAccountId(AccountId, obj, act string) bool {
	var admin *entity.CmsAdmin
	err := dao.CmsAdmin.Ctx(util.Ctx).Where(dao.CmsAdmin.Columns().Id, AccountId).Scan(&admin)
	if err != nil {
		return false
	}
	if admin == nil {
		return false
	}
	if admin.Status != 1 {
		return false
	}
	all, err := dao.CmsRoleAccount.Ctx(util.Ctx).Where("account_id", AccountId).All()
	if err != nil {
		return false
	}
	if len(all) == 0 {
		return false
	}
	var pass = false
	if len(all) == 1 {
		has, _ := initCasbin().Enforce(gvar.New(all[0]["role_id"]).String(), obj, act)
		if has {
			pass = true
		}
		one, err := dao.CmsRole.Ctx(util.Ctx).Where(dao.CmsRole.Columns().Id, all[0]["role_id"]).Where(dao.CmsRole.Columns().IsEnable, 1).One()
		if err != nil {
			return false
		}
		if one == nil {
			pass = false
		}
	} else {
		for _, one := range all {
			has, _ := initCasbin().Enforce(gvar.New(one["role_id"]).String(), obj, act)
			if has {
				pass = true
			}
		}
	}
	return pass
}

// AddByRoleId 增加权限
func (*sCasbinPolicy) AddByRoleId(roleId, obj, act string) bool {
	_, err := initCasbin().AddPolicy(roleId, obj, act)
	if err != nil {
		g.Log().Line(true).Line(true).Async(true).Async(true).Error(util.Ctx, "增加权限失败："+roleId+","+obj+","+act)
		panic(err)
	}
	return true
}

// RemoveByRoleId 删除权限
func (*sCasbinPolicy) RemoveByRoleId(roleId, obj, act string) bool {
	_, err := initCasbin().RemovePolicy(roleId, obj, act)
	if err != nil {
		g.Log().Error(util.Ctx, "删除权限失败："+roleId+","+obj+","+act)
		panic(err)
	}
	return true
}
