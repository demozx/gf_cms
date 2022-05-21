package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// Util

type sUtil struct{}

var (
	insUtil        = sUtil{}
	insViewBindFun = sViewBindFun{}
	ProjectName    *gvar.Var
	AdminPrefix    *gvar.Var
	AdminGroup     string
	AdminApiGroup  string
	Ctx            context.Context
)

func init() {
	Ctx = gctx.New()
	// 项目ProjectName
	ProjectName, _ = g.Cfg().Get(Ctx, "server.projectName", "gf_cms")
	//后台入口前缀
	AdminPrefix, _ = g.Config().Get(Ctx, "server.adminPrefix", "admin")
	//AdminGroup 后台view分组
	AdminGroup = "/" + AdminPrefix.String()
	//AdminApiGroup 后台api分组
	AdminApiGroup = "/" + AdminPrefix.String() + "_api" //system_config表根据name获取值
}

func Util() *sUtil {
	return &insUtil
}

// ProjectName 获取ProjectName
func (*sUtil) ProjectName() string {
	return ProjectName.String()
}

// AdminPrefix 后台入口前缀
func (*sUtil) AdminPrefix() string {
	return AdminPrefix.String()
}

//AdminGroup 后台view分组
func (*sUtil) AdminGroup() string {
	return "/" + Util().AdminPrefix()
}

//AdminApiGroup 后台api分组
func (*sUtil) AdminApiGroup() string {
	return "/" + Util().AdminPrefix() + "_api"
}

// GetConfig 获取配置文件的配置信息
func (*sUtil) GetConfig(node string) string {
	config, _ := g.Cfg().Get(Ctx, node)
	return config.String()
}
