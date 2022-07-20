package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/logic/setting"
	"github.com/gogf/gf/v2/frame/g"
)

type cSetting struct{}

var (
	Setting = cSetting{}
)

func (c *cSetting) Save(ctx context.Context, req *backendApi.SettingSaveApiReq) (res *backendApi.SettingSaveApiRes, err error) {
	form := g.RequestFromCtx(ctx).Request.Form
	setting.Setting().Save(form)
	g.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "保存成功",
		"data":    res,
	})
	return
}
