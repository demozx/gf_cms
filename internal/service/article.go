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
	IArticle interface {
		// MobileHomeTextNewsList 移动首页文字新闻列表
		MobileHomeTextNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error)
		MobileHomePicNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error)
		BackendArticleGetList(ctx context.Context, in *model.ArticleGetListInPut) (out *model.ArticleGetListOutPut, err error)
		Sort(ctx context.Context, in []*model.ArticleSortMap) (out interface{}, err error)
		Flag(ctx context.Context, ids []int, flagType string) (out interface{}, err error)
		Status(ctx context.Context, ids []int) (out interface{}, err error)
		Delete(ctx context.Context, ids []int) (out interface{}, err error)
		Move(ctx context.Context, channelId int, ids []string) (out interface{}, err error)
		Add(ctx context.Context, in *backendApi.ArticleAddReq) (out interface{}, err error)
		Edit(ctx context.Context, in *backendApi.ArticleEditReq) (out interface{}, err error)
		// BackendRecycleBinArticleGetList 回收站文章列表
		BackendRecycleBinArticleGetList(ctx context.Context, in *model.ArticleGetListInPut) (out *model.ArticleGetListOutPut, err error)
		// BackendRecycleBinArticleBatchDestroy 回收站-文章批量永久删除
		BackendRecycleBinArticleBatchDestroy(ctx context.Context, ids []int) (out interface{}, err error)
		// BackendRecycleBinArticleBatchRestore 回收站-文章批量恢复
		BackendRecycleBinArticleBatchRestore(ctx context.Context, ids []int) (out interface{}, err error)
		// PcHomeScrollNewsList pc首页滚动新闻
		PcHomeScrollNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error)
		// PcHomeTextNewsList pc首页文字新闻列表
		PcHomeTextNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error)
		PcHomePicNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error)
		// PcPrevArticle 上一篇文章
		PcPrevArticle(ctx context.Context, channelId int, articleId uint64) (out *model.ArticleLink, err error)
		// PcNextArticle 下一篇文章
		PcNextArticle(ctx context.Context, channelId int, articleId uint64) (out *model.ArticleLink, err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
