// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
)

type IPageInfo interface {
	LayUiPageInfo(ctx context.Context, total int, size int) string
}

var localPageInfo IPageInfo

func PageInfo() IPageInfo {
	if localPageInfo == nil {
		panic("implement not found for interface IPageInfo, forgot register?")
	}
	return localPageInfo
}

func RegisterPageInfo(i IPageInfo) {
	localPageInfo = i
}