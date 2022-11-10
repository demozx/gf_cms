package friendlyLink

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	insFriendlyLink = sFriendlyLink{}
)

type sFriendlyLink struct{}

func init() {
	service.RegisterFriendlyLink(New())
}

func New() *sFriendlyLink {
	return &sFriendlyLink{}
}

func FriendlyLink() *sFriendlyLink {
	return &insFriendlyLink
}

// BackendApiStatus 修改友情链接状态
func (s *sFriendlyLink) BackendApiStatus(ctx context.Context, req *backendApi.FriendlyLinkStatusReq) (res *backendApi.FriendlyLinkStatusRes, err error) {
	var friendlyLink *entity.CmsFriendlyLink
	err = dao.CmsFriendlyLink.Ctx(ctx).Where(dao.CmsFriendlyLink.Columns().Id, req.Id).Scan(&friendlyLink)
	if err != nil {
		return nil, err
	}
	if friendlyLink == nil {
		return nil, gerror.New("友情链接不存在")
	}
	status := 0
	if friendlyLink.Status == 0 {
		status = 1
	}
	data := g.Map{
		"id":     req.Id,
		"status": status,
	}
	_, err = dao.CmsFriendlyLink.Ctx(ctx).Where(dao.CmsFriendlyLink.Columns().Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, err
	}
	return
}
