package adminApi

import (
	"context"
	"gf_cms/api/adminApi"
	"gf_cms/internal/consts"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

func (c *cAdmin) Login(ctx context.Context, req *adminApi.AdminLoginReq) (res *adminApi.AdminLoginRes, err error) {
	admin, err := service.Admin().LoginVerify(ctx, model.AdminLoginInput{
		Username:   req.Username,
		Password:   req.Password,
		CaptchaStr: req.CaptchaStr,
		CaptchaId:  req.CaptchaId,
	})

	if err != nil {
		return
	}

	res = &adminApi.AdminLoginRes{}
	g.Dump(res)
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	g.RequestFromCtx(ctx).Session.Set(consts.AdminSessionKeyPrefix, g.Map{
		"Token":    res.Token,
		"Id":       admin.Id,
		"Username": admin.Username,
		"name":     admin.Name,
	})
	//g.Dump("token", res.Token)
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code":    0,
		"message": "登录成功",
		"data":    res,
	})
	return
}

func (c *cAdmin) Logout(ctx context.Context, req *adminApi.AdminLogoutReq) (res *adminApi.AdminLogoutRes, err error) {
	g.RequestFromCtx(ctx).Session.Remove(consts.AdminSessionKeyPrefix)
	service.Auth().LogoutHandler(ctx)
	return
}
