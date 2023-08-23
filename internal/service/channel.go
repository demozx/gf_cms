// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
)

type (
	IChannel interface {
		// HomeAboutChannel 关于我们
		HomeAboutChannel(ctx context.Context, channelId int) (channel *entity.CmsChannel, err error)
		// HomeGoodsChannelList 首页产品栏目列表
		HomeGoodsChannelList(ctx context.Context, channelId int) (out []*model.ChannelNavigationListItem, err error)
		// BackendApiIndex 获取后台栏目分类接口数据
		BackendApiIndex(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error)
		// BackendChannelTree 获取栏目分类树
		BackendChannelTree(ctx context.Context, selectedId int) (out []*model.ChannelBackendApiListItem, err error)
		BackendChannelModelTree(ctx context.Context, modelType string, channelId int) (out []*model.ChannelBackendApiListItem, err error)
		// BackendApiStatus 状态禁用启用
		BackendApiStatus(ctx context.Context, in *backendApi.ChannelStatusApiReq) (out *backendApi.ChannelStatusApiRes, err error)
		BackendApiDelete(ctx context.Context, in *backendApi.ChannelDeleteApiReq) (out *backendApi.ChannelDeleteApiRes, err error)
		BackendApiAdd(ctx context.Context, in *backendApi.ChannelAddApiReq) (out *backendApi.ChannelAddApiRes, err error)
		BackendApiEdit(ctx context.Context, in *backendApi.ChannelEditApiReq) (out *backendApi.ChannelEditApiRes, err error)
		GetOneById(ctx context.Context, id int) (out *entity.CmsChannel, err error)
		// BackendModelMap 后台数据模型映射
		BackendModelMap() map[string]string
		// BackendChannelTypeMap 后台栏目类型映射
		BackendChannelTypeMap() map[string]string
		// BackendModelCanAddMap 后台允许填充列表内容的数据模型映射
		BackendModelCanAddMap() map[string]string
		BackendModelDesc(model string) string
		BackendTypeDesc(ChannelType string) string
		// UpdateRelation 更新关联关系字段
		UpdateRelation(ctx context.Context, originChannelId int) (out interface{}, err error)
		// GetChildIds 获取当前栏目的所有子栏目id
		// belongChannelId 当前栏目
		// andMe 是否包含自身id
		GetChildIds(ctx context.Context, belongChannelId int, andMe bool) (arrAllIds []int, err error)
		// Navigation 导航
		Navigation(ctx context.Context, currChannelId int) (out []*model.ChannelNavigationListItem, err error)
		// ChildrenNavigation 获取当前栏目的子栏目
		ChildrenNavigation(ctx context.Context, navigation []*model.ChannelNavigationListItem, currChannelId int) (out []*model.ChannelNavigationListItem, err error)
		// Crumbs 生成面包屑导航
		// channelId 栏目id
		// detailId  内容页id
		Crumbs(ctx context.Context, channelId uint) (out []*model.ChannelCrumbs, err error)
		// TDK 生成pcTDK
		// channelId 栏目id
		// detailId  内容页id
		TDK(ctx context.Context, channelId uint, detailId int64) (out *model.ChannelTDK, err error)
		// MobileListTemplate 获取移动栏目列表模板
		MobileListTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error)
		// MobileDetailTemplate 获取移动栏目详情模板
		MobileDetailTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error)
		// PcListTemplate 获取pc栏目列表模板
		PcListTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error)
		// PcDetailTemplate 获取pc栏目详情模板
		PcDetailTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error)
	}
)

var (
	localChannel IChannel
)

func Channel() IChannel {
	if localChannel == nil {
		panic("implement not found for interface IChannel, forgot register?")
	}
	return localChannel
}

func RegisterChannel(i IChannel) {
	localChannel = i
}
