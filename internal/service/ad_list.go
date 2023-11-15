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
	IAdList interface {
		// Add 添加广告
		Add(ctx context.Context, req *backendApi.AdListAddReq) (out interface{}, err error)
		// Edit 编辑广告
		Edit(ctx context.Context, req *backendApi.AdListEditReq) (out interface{}, err error)
		// Delete 删除广告
		Delete(ctx context.Context, req *backendApi.AdListDeleteReq) (out interface{}, err error)
		// BatchStatus 批量修改广告状态
		BatchStatus(ctx context.Context, req *backendApi.AdListBatchStatusReq) (out interface{}, err error)
		// Sort 广告排序
		Sort(ctx context.Context, req *backendApi.AdListSortReq) (out interface{}, err error)
		// GetAdInfoById 根据广告id获取广告信息
		GetAdInfoById(ctx context.Context, id int) (out interface{}, err error)
		// PcHomeListByChannelId pc首页banner
		PcHomeListByChannelId(ctx context.Context, channelId int) (out []*entity.CmsAd, err error)
	}
)

var (
	localAdList IAdList
)

func AdList() IAdList {
	if localAdList == nil {
		panic("implement not found for interface IAdList, forgot register?")
	}
	return localAdList
}

func RegisterAdList(i IAdList) {
	localAdList = i
}
