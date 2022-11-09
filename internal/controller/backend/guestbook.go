package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Guestbook = cGuestbook{}
)

type cGuestbook struct{}

func (c *cGuestbook) Index(ctx context.Context, req *backend.GuestbookIndexReq) (res *backend.GuestbookIndexRes, err error) {
	list, err := service.Guestbook().BackendGetList(ctx, req)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/guestbook/index.html", g.Map{
		"list":     list,
		"pageInfo": service.PageInfo().LayUiPageInfo(ctx, list.Total, list.Size),
	})
	if err != nil {
		return nil, err
	}
	return
}
