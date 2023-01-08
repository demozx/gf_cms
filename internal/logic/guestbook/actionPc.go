package guestbook

import (
	"context"
	"gf_cms/api/pcApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcSubmit pc提交留言
func (s *sGuestbook) PcSubmit(ctx context.Context, in *pcApi.GuestbookReq) (out *pcApi.GuestbookRes, err error) {
	var guestbook *entity.CmsGuestbook
	err = dao.CmsGuestbook.Ctx(ctx).Where(dao.CmsGuestbook.Columns().Tel, in.Tel).Where(dao.CmsGuestbook.Columns().Status, 0).Scan(&guestbook)
	if err != nil {
		return nil, err
	}
	if guestbook != nil {
		return nil, gerror.New("留言提交失败，您已于<br/>" + gconv.String(guestbook.CreatedAt) + "<br/>提交过留言")
	}
	ip := g.RequestFromCtx(ctx).GetClientIp()
	data := g.Map{
		dao.CmsGuestbook.Columns().Name:    in.Name,
		dao.CmsGuestbook.Columns().Tel:     in.Tel,
		dao.CmsGuestbook.Columns().Content: in.Email + "<br>" + in.Content,
		dao.CmsGuestbook.Columns().Ip:      ip,
		dao.CmsGuestbook.Columns().From:    1,
	}
	lastId, err := dao.CmsGuestbook.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	// 更新ip归属地
	go func() {
		g.Log().Debug(ctx, "留言：更新ip归属地")
		address, err := service.Guestbook().GetAddressByIp(ctx, ip)
		_, err = dao.CmsGuestbook.Ctx(ctx).Where(dao.CmsGuestbook.Columns().Id, lastId).Data(g.Map{
			dao.CmsGuestbook.Columns().Address: address,
		}).Update()
		if err != nil {
			return
		}
		// 发送邮件
		_, err = service.Guestbook().SendEmail(ctx, lastId)
		if err != nil {
			return
		}
	}()
	return
}
