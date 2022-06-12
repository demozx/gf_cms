package backend

import (
	"context"
	"gf_cms/api/backend"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Setting = cSetting{}
)

type cSetting struct{}

// Index 后台设置
func (c *cSetting) Index(ctx context.Context, req *backend.SettingReq) (res *backend.SettingRes, err error) {
	_ = g.RequestFromCtx(ctx).Response.WriteTpl("backend/setting/index.html", g.Map{})
	return
}
