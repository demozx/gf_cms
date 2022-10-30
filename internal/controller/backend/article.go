package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

func (c *cArticle) Move(ctx context.Context, req *backend.ArticleMoveReq) (res *backend.ArticleMoveRes, err error) {
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelArticle, 0)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/article/move.html", g.Map{
		"strIds":           req.StrIds,
		"channelModelTree": channelModelTree,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cArticle) Add(ctx context.Context, req *backend.ArticleAddReq) (res *backend.ArticleAddRes, err error) {
	channelId := req.ChannelId
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelArticle, channelId)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/article/add.html", g.Map{
		"channelModelTree": channelModelTree,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cArticle) Edit(ctx context.Context, req *backend.ArticleEditReq) (res *backend.ArticleEditRes, err error) {
	var articleInfo *model.ArticleWithBody
	err = dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, req.Id).WithAll().Scan(&articleInfo)
	if err != nil {
		return nil, err
	}
	if articleInfo == nil {
		return nil, gerror.New("文章不存在")
	}
	// 将flag变成数组，方便模板渲染
	flagArr := gstr.Explode(",", articleInfo.Flag)
	for _, value := range flagArr {
		if value == "p" {
			articleInfo.FlagP = 1
		}
		if value == "r" {
			articleInfo.FlagR = 1
		}
		if value == "t" {
			articleInfo.FlagT = 1
		}
	}
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelArticle, articleInfo.ChannelId)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/article/edit.html", g.Map{
		"channelModelTree": channelModelTree,
		"articleInfo":      articleInfo,
	})
	if err != nil {
		return nil, err
	}
	return
}
