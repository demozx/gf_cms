// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
)

type (
	IImage interface {
		// PcPrevImage 上一篇图集
		PcPrevImage(ctx context.Context, channelId int, imageId uint64) (out *model.ImageLink, err error)
		// PcNextImage 下一篇图集
		PcNextImage(ctx context.Context, channelId int, imageId uint64) (out *model.ImageLink, err error)
		BackendImageGetList(ctx context.Context, in *model.ImageGetListInPut) (out *model.ImageGetListOutPut, err error)
		Sort(ctx context.Context, in []*model.ImageSortMap) (out interface{}, err error)
		Flag(ctx context.Context, ids []int, flagType string) (out interface{}, err error)
		Status(ctx context.Context, ids []int) (out interface{}, err error)
		Delete(ctx context.Context, ids []int) (out interface{}, err error)
		Move(ctx context.Context, channelId int, ids []string) (out interface{}, err error)
		Add(ctx context.Context, in *backendApi.ImageAddReq) (out interface{}, err error)
		Edit(ctx context.Context, in *backendApi.ImageEditReq) (out interface{}, err error)
		// BackendRecycleBinImageGetList 回收站图集列表
		BackendRecycleBinImageGetList(ctx context.Context, in *model.ImageGetListInPut) (out *model.ImageGetListOutPut, err error)
		// BackendRecycleBinImageBatchDestroy 回收站-图集批量永久删除
		BackendRecycleBinImageBatchDestroy(ctx context.Context, ids []int) (out interface{}, err error)
		// BackendRecycleBinImageBatchRestore 回收站-图集批量恢复
		BackendRecycleBinImageBatchRestore(ctx context.Context, ids []int) (out interface{}, err error)
		// BuildThumb 构建图集缩略图
		BuildThumb(ctx context.Context, in *model.ImageListItem) (out *model.ImageListItem, err error)
		// MobileHomeRecommendGoodsList 移动首页推荐产品图集列表
		MobileHomeRecommendGoodsList(ctx context.Context, belongChannelId int) (out []*model.ImageListItem, err error)
		// MobileHomeHonerList 移动首页荣誉资质列表
		MobileHomeHonerList(ctx context.Context, channelId int) (out []*model.ImageListItem, err error)
		// PcHomeRecommendGoodsList pc首页推荐产品图集列表
		PcHomeRecommendGoodsList(ctx context.Context, belongChannelId int) (out []*model.ImageListItem, err error)
		PcHomeGoodsGroupList(ctx context.Context, belongChannelId int) (out [][]*model.ImageListItem, err error)
	}
)

var (
	localImage IImage
)

func Image() IImage {
	if localImage == nil {
		panic("implement not found for interface IImage, forgot register?")
	}
	return localImage
}

func RegisterImage(i IImage) {
	localImage = i
}
