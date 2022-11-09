package guestbook

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	insGuestbook = sGuestbook{}
)

type sGuestbook struct{}

func init() {
	service.RegisterGuestbook(New())
}

func New() *sGuestbook {
	return &sGuestbook{}
}

func Menu() *sGuestbook {
	return &insGuestbook
}

// BackendGetList 留言板列表
func (s *sGuestbook) BackendGetList(ctx context.Context, in *backend.GuestbookIndexReq) (out *model.GuestbookGetListOutput, err error) {
	var list []*model.GuestbookGetListOutputItem
	m := dao.CmsGuestbook.Ctx(ctx).OrderAsc(dao.CmsGuestbook.Columns().Status).OrderDesc(dao.CmsGuestbook.Columns().Id)
	err = m.Page(in.Page, in.Size).Scan(&list)
	if err != nil {
		return nil, err
	}
	for key, item := range list {
		fromDesc := ""
		if item.From == 1 {
			fromDesc = "电脑端"
		} else if item.From == 2 {
			fromDesc = "移动端"
		}
		list[key].FromDesc = fromDesc
	}
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	out = &model.GuestbookGetListOutput{
		List:  list,
		Total: count,
		Page:  in.Page,
		Size:  in.Size,
	}
	return
}

// BackendApiStatus 修改留言状态
func (s *sGuestbook) BackendApiStatus(ctx context.Context, in *backendApi.GuestbookStatusReq) (out interface{}, err error) {
	var guestbook *entity.CmsGuestbook
	err = dao.CmsGuestbook.Ctx(ctx).Where(dao.CmsGuestbook.Columns().Id, in.Id).Scan(&guestbook)
	if err != nil {
		return nil, err
	}
	if guestbook == nil {
		return nil, gerror.New("留言不存在")
	}
	status := 0
	if guestbook.Status == 0 {
		status = 1
	}
	data := g.Map{
		dao.CmsGuestbook.Columns().Status: status,
	}
	_, err = dao.CmsGuestbook.Ctx(ctx).Where(dao.CmsGuestbook.Columns().Id, in.Id).Data(data).Update()
	if err != nil {
		return nil, err
	}
	return
}

// BackendApiBatchDelete 批量删除留言
func (s *sGuestbook) BackendApiBatchDelete(ctx context.Context, in *backendApi.GuestbookDeleteReq) (out interface{}, err error) {
	_, err = dao.CmsGuestbook.Ctx(ctx).WhereIn(dao.CmsGuestbook.Columns().Id, in.Ids).Delete()
	if err != nil {
		return nil, err
	}
	return
}
