package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/admin"
	"gf_cms/internal/logic/auth"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

func (c *cAdmin) Login(ctx context.Context, req *backendApi.AdminLoginReq) (res *backendApi.AdminLoginRes, err error) {
	admin, err := admin.Admin().LoginVerify(ctx, model.AdminLoginInput{
		Username:   req.Username,
		Password:   req.Password,
		CaptchaStr: req.CaptchaStr,
		CaptchaId:  req.CaptchaId,
	})

	if err != nil {
		return
	}

	res = &backendApi.AdminLoginRes{}
	//生成token
	res.Token, res.Expire = auth.Auth().JWTAuth().LoginHandler(ctx)
	// 记录session：自己定义的，因为一般后台登录用session
	g.RequestFromCtx(ctx).Session.Set(consts.AdminSessionKeyPrefix, g.Map{
		"Token":    res.Token,
		"Id":       admin.Id,
		"Username": admin.Username,
		"name":     admin.Name,
	})
	g.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "登录成功",
		"data":    res,
	})
	return
}

func (c *cAdmin) Logout(ctx context.Context, req *backendApi.AdminLogoutReq) (res *backendApi.AdminLogoutRes, err error) {
	//清除session
	g.RequestFromCtx(ctx).Session.Remove(consts.AdminSessionKeyPrefix)
	//清除token
	auth.Auth().JWTAuth().LogoutHandler(ctx)
	return
}

func (c *cAdmin) ClearCache(ctx context.Context, req *backendApi.AdminClearCacheReq) (res *backendApi.AdminClearCacheRes, err error) {
	_, err = util.Util().ClearPublicCache()
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
			"code":    0,
			"message": "清理成功",
		})
	}
	return
}
