// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

type (
	IUtil interface {
		// DryRun 获取空跑模式
		DryRun() bool
		// ProjectName 获取ProjectName
		ProjectName() string
		// SystemRoot 获取SystemRoot
		SystemRoot() string
		// BackendPrefix 后台入口前缀
		BackendPrefix() string
		// BackendApiPrefix 后台入口前缀
		BackendApiPrefix() string
		// BackendGroup 后台view分组
		BackendGroup() string
		// BackendApiGroup 后台api分组
		BackendApiGroup() string
		// PcApiGroup pcApi分组
		PcApiGroup() string
		// MApiGroup 移动Api分组
		MApiGroup() string
		// ServerRoot 服务目录
		ServerRoot() string
		// GetConfig 获取配置文件的配置信息
		GetConfig(node string) string
		// GetSetting 获取设置
		GetSetting(name string) string
		// ClearPublicCache 清除公共缓存
		ClearPublicCache() (*gvar.Var, error)
		// ClearSystemSettingCache 清除后台设置缓存
		ClearSystemSettingCache() (*gvar.Var, error)
		// GetLocalIP 获取ip
		GetLocalIP() (ip string, err error)
		// FriendyTimeFormat 计算时间差，并以"XXd XXh XXm XXs"返回
		FriendyTimeFormat(TimeCreate time.Time, TimeEnd time.Time) string
		// ImageOrDefaultUrl 返回图片或默认图url
		ImageOrDefaultUrl(imgUrl string) string
		// IsMobile 判断是手机端
		IsMobile(ctx context.Context) bool
		// ResponsiveJump 响应跳转
		ResponsiveJump(ctx context.Context)
	}
)

var (
	localUtil IUtil
)

func Util() IUtil {
	if localUtil == nil {
		panic("implement not found for interface IUtil, forgot register?")
	}
	return localUtil
}

func RegisterUtil(i IUtil) {
	localUtil = i
}
