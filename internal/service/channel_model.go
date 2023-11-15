// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
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
		// ModelArticle 文章模型列表
		ModelArticle(ctx context.Context, in *backend.ChannelModelIndexReq) (out []*model.ChannelBackendApiListItem, err error)
		// ModelImage 图集模型列表
		ModelImage(ctx context.Context, in *backend.ChannelModelIndexReq) (out []*model.ChannelBackendApiListItem, err error)
		// GetDetailOneByChannelId 根据栏目id，详情页id，获取一条对应的详情
		GetDetailOneByChannelId(ctx context.Context, channelId uint, detailId int64) (out interface{}, err error)
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
