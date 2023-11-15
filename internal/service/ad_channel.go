// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model/entity"
)

type (
	IAdChannel interface {
		Add(ctx context.Context, in *backendApi.AdChannelAddReq) (out interface{}, err error)
		Edit(ctx context.Context, in *backendApi.AdChannelEditReq) (out interface{}, err error)
		Delete(ctx context.Context, in *backendApi.AdChannelDeleteReq) (out interface{}, err error)
		Sort(ctx context.Context, in *backendApi.AdChannelSortReq) (out interface{}, err error)
		// GetAdChannelMap 获取adChannelMap
		GetAdChannelMap(ctx context.Context) (output []*entity.CmsAdChannel, err error)
	}
)

var (
	localAdChannel IAdChannel
)

func AdChannel() IAdChannel {
	if localAdChannel == nil {
		panic("implement not found for interface IAdChannel, forgot register?")
	}
	return localAdChannel
}

func RegisterAdChannel(i IAdChannel) {
	localAdChannel = i
}
