package viewBindFun

import (
	"gf_cms/internal/logic/util"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	insViewBindFun = sViewBindFun{}
)

// ViewBindFun 模板视图绑定的方法
func ViewBindFun() *sViewBindFun {
	return &insViewBindFun
}

type sViewBindFun struct{}

func init() {
	service.RegisterViewBindFun(New())
}

func New() *sViewBindFun {
	return &sViewBindFun{}
}

// Register 注册视图绑定方法
func (*sViewBindFun) Register() {
	g.View().BindFuncMap(g.Map{
		"system_setting":  ViewBindFun().SystemSetting,
		"system_config":   ViewBindFun().SystemConfig,
		"backend_url":     ViewBindFun().BackendUrl,
		"backend_api_url": ViewBindFun().BackendApiUrl,
		"pc_api_url":      ViewBindFun().PcApiUrl,
		"mobile_api_url":  ViewBindFun().MobileApiUrl,
		"dry_run":         ViewBindFun().DryRun,
	})
}

// SystemSetting 获取系统配置信息
func (*sViewBindFun) SystemSetting(name string) string {
	return service.Util().GetSetting(name)
}

// SystemConfig 获取系统配置
func (*sViewBindFun) SystemConfig(name string) string {
	return service.Util().GetConfig(name)
}

// BackendUrl 生成后台view的url
func (*sViewBindFun) BackendUrl(route string) string {
	return util.BackendGroup + route
}

// BackendApiUrl 生成后台api的url
func (*sViewBindFun) BackendApiUrl(route string) string {
	return util.BackendApiGroup + route
}

// PcApiUrl 生成pcApi发的路由
func (*sViewBindFun) PcApiUrl(route string) string {
	return util.PcApiGroup + route
}

// MobileApiUrl 生成MApi发的路由
func (*sViewBindFun) MobileApiUrl(route string) string {
	return util.MobileApiGroup + route
}

func (*sViewBindFun) DryRun() bool {
	return util.DryRun
}
