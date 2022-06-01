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
	insUtil           = sUtil{}
	insViewBindFun    = sViewBindFun{}
	ProjectName       *gvar.Var
	BackendPrefix     *gvar.Var
	BackendGroup      string
	BackendApiGroup   string
	Ctx               context.Context
	PublicCachePreFix string
)

func init() {
	Ctx = gctx.New()
	//项目ProjectName
	ProjectName, _ = g.Cfg().Get(Ctx, "server.projectName", "gf_cms")
	//后台入口前缀
	BackendPrefix, _ = g.Config().Get(Ctx, "server.backendPrefix", "admin")
	//BackendGroup 后台view分组
	BackendGroup = "/" + BackendPrefix.String()
	//BackendApiGroup 后台api分组
	BackendApiGroup = "/" + BackendPrefix.String() + "_api"
	//公共缓存前缀
	PublicCachePreFix = ProjectName.String() + ":public"
}

func Util() *sUtil {
	return &insUtil
}

// ProjectName 获取ProjectName
func (*sUtil) ProjectName() string {
	return ProjectName.String()
}

// BackendPrefix 后台入口前缀
func (*sUtil) BackendPrefix() string {
	return BackendPrefix.String()
}

//BackendGroup 后台view分组
func (*sUtil) BackendGroup() string {
	return "/" + Util().BackendPrefix()
}

//BackendApiGroup 后台api分组
func (*sUtil) BackendApiGroup() string {
	return "/" + Util().BackendPrefix() + "_api"
}

// GetConfig 获取配置文件的配置信息
func (*sUtil) GetConfig(node string) string {
	config, _ := g.Cfg().Get(Ctx, node)
	return config.String()
}

// ClearPublicCache 清除公共缓存
func (*sUtil) ClearPublicCache() (*gvar.Var, error) {
	cacheKey := PublicCachePreFix + ":*"
	keys, err := g.Redis().Do(Ctx, "KEYS", cacheKey)
	if err != nil {
		return nil, err
	}
	for _, key := range keys.Array() {
		_, err := g.Redis().Do(Ctx, "DEL", key)
		if err != nil {
			return nil, err
		}
	}
	return nil, err
}
