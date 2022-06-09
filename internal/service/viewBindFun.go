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
		"system_setting":  ViewBindFun().SystemSetting,
		"backend_url":     ViewBindFun().BackendUrl,
		"backend_api_url": ViewBindFun().BackendApiUrl,
	})
}

// SystemSetting 获取系统配置信息
func (*sViewBindFun) SystemSetting(name string) string {
	cacheKey := PublicCachePreFix + ":system_setting:" + name
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
	val, _ := g.Model("system_setting").Where("name", name).Value("value")
	g.Redis().Do(Ctx, "SET", cacheKey, val.String())
	return val.String()
}

// BackendUrl 生成后台view的url
func (*sViewBindFun) BackendUrl(route string) string {
	return BackendGroup + route
}

// BackendApiUrl 生成后台api的url
func (*sViewBindFun) BackendApiUrl(route string) string {
	return BackendApiGroup + route
}
