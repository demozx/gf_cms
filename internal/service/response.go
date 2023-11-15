// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IResponse interface {
		// SuccessJson 返回成功json
		SuccessJson(ctx context.Context, code int, message string, data interface{})
		// SuccessJsonDefault 返回默认成功json
		SuccessJsonDefault(ctx context.Context)
		// SuccessJsonDefaultMessage 返回自定义提示语的成功json
		SuccessJsonDefaultMessage(ctx context.Context, message string)
		// SuccessCodeDefault 获取默认成功code码
		SuccessCodeDefault() int
		// SuccessMessageDefault 获取默认成功提示语
		SuccessMessageDefault() string
		// View 模板渲染
		View(ctx context.Context, template string, data map[string]interface{}) (err error)
		// ErrorTpl 错误页面
		ErrorTpl(ctx context.Context, code int, message string) (err error)
		// Status404 状态码404
		Status404(ctx context.Context)
		// Status500 http状态码500
		Status500(ctx context.Context)
	}
)

var (
	localResponse IResponse
)

func Response() IResponse {
	if localResponse == nil {
		panic("implement not found for interface IResponse, forgot register?")
	}
	return localResponse
}

func RegisterResponse(i IResponse) {
	localResponse = i
}
