package adminApi

import (
	"context"
	"gf_cms/api/adminApi"
	"gf_cms/internal/service"
)

var (
	Captcha = cCaptcha{}
)

type cCaptcha struct{}

func (c *cCaptcha) Get(ctx context.Context, req *adminApi.CaptchaGetApiReq) (res *adminApi.CaptchaGetApiRes, err error) {
	id, b64s := service.Captcha().Get()
	res = &adminApi.CaptchaGetApiRes{
		Id: id, B64s: b64s,
	}

	return
}
