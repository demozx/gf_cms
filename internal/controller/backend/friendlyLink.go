package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	FriendlyLink = cFriendlyLink{}
)

type cFriendlyLink struct{}

// Index 友情链接列表
func (c *cFriendlyLink) Index(ctx context.Context, req *backend.FriendlyLinkIndexReq) (res *backend.FriendlyLinkIndexRes, err error) {
	var friendlyLinkList []*entity.CmsFriendlyLink
	err = dao.CmsFriendlyLink.Ctx(ctx).OrderAsc(dao.CmsFriendlyLink.Columns().Sort).OrderAsc(dao.CmsFriendlyLink.Columns().Id).Scan(&friendlyLinkList)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/friendly_link/index.html", g.Map{
		"list":  friendlyLinkList,
		"total": len(friendlyLinkList),
	})
	if err != nil {
		return nil, err
	}
	return
}
