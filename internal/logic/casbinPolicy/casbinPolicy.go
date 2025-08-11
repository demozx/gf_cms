package casbinPolicy

import (
	"database/sql"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/casbin/casbin/v2/model"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
	"time"

	sqlAdapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	insPolicy            = sCasbinPolicy{}
	ObjBackend    string = "backend"
	ObjBackendApi string = "backend_api"
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
	var dbType = util.Util().GetConfig("database.default.type")
	var dbUser = util.Util().GetConfig("database.default.user")
	var dbPass = util.Util().GetConfig("database.default.pass")
	var dbHost = util.Util().GetConfig("database.default.host")
	var dbPort = util.Util().GetConfig("database.default.port")
	var dbName = util.Util().GetConfig("database.default.name")
	var dbPrefix = util.Util().GetConfig("database.default.prefix")
	db, err := sql.Open(dbType, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(gconv.Int(util.Util().GetConfig("database.default.maxOpen")))
	db.SetMaxIdleConns(gconv.Int(util.Util().GetConfig("database.default.maxIdle")))
	db.SetConnMaxLifetime(time.Second * 10)
	a, err := sqlAdapter.NewAdapter(db, dbType, dbPrefix+"rule_permissions")
	if err != nil {
		panic(err)
	}
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	if err != nil {
		panic(err)
	}
	// Load the policy from DB.
	if err = e.LoadPolicy(); err != nil {
		log.Println("LoadPolicy failed, err: ", err)
	}
	//defer db.Close()
	db.Close()
	return e
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
