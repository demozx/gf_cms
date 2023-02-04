package guestbook

import (
	"context"
	"gf_cms/api/mobileApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// MobileSubmit mobile提交留言
func (s *sGuestbook) MobileSubmit(ctx context.Context, in *mobileApi.GuestbookReq) (out *mobileApi.GuestbookRes, err error) {
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
		dao.CmsGuestbook.Columns().From:    2,
	}
	lastId, err := dao.CmsGuestbook.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		glog.Error(ctx, err)
		return nil, gerror.New("提交留言出错")
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
