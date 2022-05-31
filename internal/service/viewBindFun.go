package service

import (
	"github.com/gogf/gf/v2/frame/g"
)

//模板视图绑定的方法

func ViewBindFun() *sViewBindFun {
	return &insViewBindFun
}

type sViewBindFun struct{}

// Register 注册视图绑定方法
func (*sViewBindFun) Register() {
	g.View().BindFuncMap(g.Map{
		"system_config": ViewBindFun().SystemConfig,
		"admin_url":     ViewBindFun().AdminUrl,
		"admin_api_url": ViewBindFun().AdminApiUrl,
	})
}

// SystemConfig 获取系统配置信息
func (*sViewBindFun) SystemConfig(name string) string {
	cacheKey := PublicCachePreFix + ":system_config:" + name
	exists, err := g.Redis().Do(Ctx, "EXISTS", cacheKey)
	if err != nil {
		panic(err)
	}
	//存在缓存key
	if exists.Bool() {
		value, err := g.Redis().Do(Ctx, "GET", cacheKey)
		if err != nil {
			panic(err)
		}
		return value.String()
	}
	//不存在缓存key，从数据库读取
	val, _ := g.Model("system_config").Where("name", name).Value("value")
	g.Redis().Do(Ctx, "SET", cacheKey, val.String())
	return val.String()
}

// AdminUrl 生成后台view的url
func (*sViewBindFun) AdminUrl(route string) string {
	return BackendGroup + route
}

// AdminApiUrl 生成后台api的url
func (*sViewBindFun) AdminApiUrl(route string) string {
	return BackendApiGroup + route
}
