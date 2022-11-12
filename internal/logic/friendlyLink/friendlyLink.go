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
func (s *sFriendlyLink) BackendApiStatus(ctx context.Context, req *backendApi.FriendlyLinkStatusReq) (res interface{}, err error) {
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

// BackendApiAdd 添加友情链接
func (s *sFriendlyLink) BackendApiAdd(ctx context.Context, req *backendApi.FriendlyLinkAddReq) (res interface{}, err error) {
	var cmsFriendlyLink *entity.CmsFriendlyLink
	err = dao.CmsFriendlyLink.Ctx(ctx).Where(dao.CmsFriendlyLink.Columns().Url, req.Url).Scan(&cmsFriendlyLink)
	if err != nil {
		return nil, err
	}
	if cmsFriendlyLink != nil {
		return nil, gerror.New("链接地址已存在")
	}
	_, err = dao.CmsFriendlyLink.Ctx(ctx).Data(req).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sFriendlyLink) BackendApiEdit(ctx context.Context, req *backendApi.FriendlyLinkEditReq) (res interface{}, err error) {
	var cmsFriendlyLink *entity.CmsFriendlyLink
	err = dao.CmsFriendlyLink.Ctx(ctx).Where(dao.CmsFriendlyLink.Columns().Url, req.Url).WhereNot(dao.CmsFriendlyLink.Columns().Id, req.Id).Scan(&cmsFriendlyLink)
	if err != nil {
		return nil, err
	}
	if cmsFriendlyLink != nil {
		return nil, gerror.New("链接地址已存在")
	}
	affected, err := dao.CmsFriendlyLink.Ctx(ctx).Where(dao.CmsFriendlyLink.Columns().Id, req.Id).Data(req).UpdateAndGetAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, gerror.New("友情链接不存在")
	}
	return
}
