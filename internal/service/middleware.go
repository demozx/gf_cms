// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		CORS(r *ghttp.Request)
		Auth(r *ghttp.Request)
		BackendAuthSession(r *ghttp.Request)
		// BackendCheckPolicy 检测后台页面用户有无某个请求权限
		BackendCheckPolicy(r *ghttp.Request)
		PcResponse(r *ghttp.Request)
		MobileResponse(r *ghttp.Request)
		// BackendApiCheckPolicy 检测后台接口用户有无某个请求权限
		BackendApiCheckPolicy(r *ghttp.Request)
		// GetBackendUserID 获取后台用户ID
		GetBackendUserID(r *ghttp.Request) (accountId string, err error)
		// FilterXSS 过滤xss攻击
		FilterXSS(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
