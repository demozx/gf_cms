package guestbook

import (
	"context"
	"crypto/tls"
	"gf_cms/api/pcApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"gopkg.in/gomail.v2"
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
		dao.CmsGuestbook.Columns().Content: in.Email + "\n\r" + in.Content,
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

// GetAddressByIp 根据ip获取归属地
func (s *sGuestbook) GetAddressByIp(ctx context.Context, ip string) (address string, err error) {
	content := g.Client().GetContent(ctx, "http://api.map.baidu.com/location/ip?ip="+ip+"&ak=8139d029b1608c4f6f6ff182318debb5&coor=bd09ll")
	if err != nil {
		return "", err
	}
	var baiduAddressByIp *model.BaiduAddressByIp
	err = gconv.Scan(content, &baiduAddressByIp)
	if err != nil {
		return "", err
	}
	if baiduAddressByIp.Status != 0 {
		return "", gerror.New(baiduAddressByIp.Message)
	}
	address = baiduAddressByIp.Content.Address
	return
}

// SendEmail 发送留言邮件
func (s *sGuestbook) SendEmail(ctx context.Context, guestbookId int64) (out interface{}, err error) {
	g.Log().Debug(ctx, "留言：邮件提醒")
	emailNotice := service.Util().GetSetting("guestbook_email_notice")
	adminEmails := service.Util().GetSetting("admin_emails")
	if emailNotice != "1" {
		g.Log().Error(ctx, "留言邮件提醒管理员未开启")
		return
	}
	if adminEmails == "" {
		g.Log().Error(ctx, "管理员邮箱未填写")
		return
	}
	var guestbook *entity.CmsGuestbook
	err = dao.CmsGuestbook.Ctx(ctx).Where(dao.CmsGuestbook.Columns().Id, guestbookId).Scan(&guestbook)
	if err != nil {
		return nil, err
	}
	host := service.Util().GetSetting("smtp_server")
	port := service.Util().GetSetting("smtp_port")
	from := service.Util().GetSetting("smtp_email_from")
	password := service.Util().GetSetting("smtp_pass")
	if host == "" || port == "" || from == "" || password == "" {
		g.Log().Error(ctx, "邮件服务器配置缺失")
		return
	}
	to := adminEmails
	subject := "留言提醒"
	guestbookFrom := ""
	switch guestbook.From {
	case 1:
		guestbookFrom = "电脑端"
	case 2:
		guestbookFrom = "移动端"
	}
	body := "网站有新留言，详细内容如下：<br><br>"
	body += "姓名：" + guestbook.Name + "<br>"
	body += "手机：" + guestbook.Tel + "<br>"
	body += "留言内容：" + guestbook.Content + "<br>"
	body += "来源：" + guestbookFrom + "<br>"
	body += "IP：" + guestbook.Ip + "<br>"
	body += "归属地：" + guestbook.Address + "<br>"
	body += "<br>以上内容已经同步到网站后台，请登录后进行处理"
	m := gomail.NewMessage()
	m.SetHeader(`From`, from)
	m.SetHeader(`To`, to)
	m.SetHeader(`Subject`, subject)
	m.SetBody("text/html", body)
	// 下面的配置改成你自己的邮箱配置
	d := gomail.NewDialer(host, gconv.Int(port), from, password)
	// 修改TLSconfig
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err = d.DialAndSend(m); err != nil {
		return
	}
	return
}
