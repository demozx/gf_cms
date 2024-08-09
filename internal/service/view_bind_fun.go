// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IViewBindFun interface {
		// Register 注册视图绑定方法
		Register()
		// SystemSetting 获取系统配置信息
		SystemSetting(name string) (setting string, err error)
		// SystemConfig 获取系统配置
		SystemConfig(name string) string
		// BackendUrl 生成后台view的url
		BackendUrl(route string) string
		// BackendApiUrl 生成后台api的url
		BackendApiUrl(route string) string
		// PcApiUrl 生成pcApi发的路由
		PcApiUrl(route string) string
		// MobileApiUrl 生成MApi发的路由
		MobileApiUrl(route string) string
		DryRun() bool
	}
)

var (
	localViewBindFun IViewBindFun
)

func ViewBindFun() IViewBindFun {
	if localViewBindFun == nil {
		panic("implement not found for interface IViewBindFun, forgot register?")
	}
	return localViewBindFun
}

func RegisterViewBindFun(i IViewBindFun) {
	localViewBindFun = i
}
