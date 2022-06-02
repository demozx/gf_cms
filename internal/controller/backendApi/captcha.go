package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

var (
	Captcha = cCaptcha{}
)

type cCaptcha struct{}

func (c *cCaptcha) Get(ctx context.Context, req *backendApi.CaptchaGetApiReq) (res *backendApi.CaptchaGetApiRes, err error) {
	id, b64s := service.Captcha().Get()
	res = &backendApi.CaptchaGetApiRes{
		Id: id, B64s: b64s,
	}

	return
}
