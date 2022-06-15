package service

import (
	"database/sql"
	"gf_cms/internal/service/internal/dao"
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"log"
)

var (
	insPolicy         = sPolicy{}
	ObjBackend string = "backend"
)

type sPolicy struct{}

func CasbinPolicy() *sPolicy {
	return &insPolicy
}

//初始化casbin
func initCasbin() *casbin.Enforcer {
	var dbType = Util().GetConfig("database.default.type")
	var dbUser = Util().GetConfig("database.default.user")
	var dbPass = Util().GetConfig("database.default.pass")
	var dbHost = Util().GetConfig("database.default.host")
	var dbPort = Util().GetConfig("database.default.port")
	var dbName = Util().GetConfig("database.default.name")
	var dbPrefix = Util().GetConfig("database.default.prefix")
	db, err := sql.Open(dbType, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	//db.SetMaxOpenConns(20)
	//db.SetMaxIdleConns(10)
	//db.SetConnMaxLifetime(time.Minute * 10)
	a, err := sqladapter.NewAdapter(db, dbType, dbPrefix+"rule_permissions")
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer(Util().SystemRoot()+"/manifest/config/rbac_model.conf", a)
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
	defer db.Close()
	return e
}

// ObjBackend 获取后台obj
func (*sPolicy) ObjBackend() string {
	return ObjBackend
}

// CheckByRoleId 检测角色权限
func (*sPolicy) CheckByRoleId(roleId, obj, act string) bool {
	has, err := initCasbin().Enforce(roleId, obj, act)
	if err != nil {
		log.Println("Enforce failed, err: ", err)
		g.Log().Error(Ctx, "Enforce failed, err: ", err)
	}
	if !has {
		g.Log().Warning(Ctx, "没有操作权限："+roleId+","+obj+","+act)
	} else {
		return true
	}
	return false
}

// CheckByAccountId 检测用户权限
func (*sPolicy) CheckByAccountId(AccountId, obj, act string) bool {
	all, err := dao.CmsRoleAccount.Ctx(Ctx).Where("account_id", AccountId).All()
	if err != nil {
		return false
	}
	var pass = false
	for _, one := range all {
		has, _ := initCasbin().Enforce(gvar.New(one["role_id"]).String(), obj, act)
		if has {
			pass = true
		}
	}

	return pass
}

// AddByRoleId 增加权限
func (*sPolicy) AddByRoleId(roleId, obj, act string) bool {
	_, err := initCasbin().AddPolicy(roleId, obj, act)
	if err != nil {
		g.Log().Error(Ctx, "增加权限失败："+roleId+","+obj+","+act)
		panic(err)
	}
	return true
}

// RemoveByRoleId 删除权限
func (*sPolicy) RemoveByRoleId(roleId, obj, act string) bool {
	_, err := initCasbin().RemovePolicy(roleId, obj, act)
	if err != nil {
		g.Log().Error(Ctx, "删除权限失败："+roleId+","+obj+","+act)
		panic(err)
	}
	return true
}
