package guestbook

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
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
