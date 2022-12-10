// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IGenUrl interface {
		PcChannelUrl(ctx context.Context, channelId int, router string) (newRouter string, err error)
		PcDetailUrl(ctx context.Context, model string, detailId int) (newRouter string, err error)
	}
)

var (
	localGenUrl IGenUrl
)

func GenUrl() IGenUrl {
	if localGenUrl == nil {
		panic("implement not found for interface IGenUrl, forgot register?")
	}
	return localGenUrl
}

func RegisterGenUrl(i IGenUrl) {
	localGenUrl = i
}