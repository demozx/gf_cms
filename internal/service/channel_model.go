// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/model"
)

type (
	IChannelModel interface {
		ModelArticle(ctx context.Context, in *backend.ChannelModelIndexReq) (out []*model.ChannelBackendApiListItem, err error)
	}
)

var (
	localChannelModel IChannelModel
)

func ChannelModel() IChannelModel {
	if localChannelModel == nil {
		panic("implement not found for interface IChannelModel, forgot register?")
	}
	return localChannelModel
}

func RegisterChannelModel(i IChannelModel) {
	localChannelModel = i
}