package pageInfo

import (
	"context"
	"gf_cms/internal/service"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type sPageInfo struct {
}

var pageInfo = sPageInfo{}

func init() {
	service.RegisterPageInfo(New())
}

func New() *sPageInfo {
	return &sPageInfo{}
}

func PageInfo() *sPageInfo {
	return &pageInfo
}

// LayUiPageInfo layui分页
func (s *sPageInfo) LayUiPageInfo(ctx context.Context, total int, size int) string {
	page := g.RequestFromCtx(ctx).GetPage(total, size)
	pageInfo := page.GetContent(3)
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<span class=\"GPageSpan\">" + gvar.New(page.CurrentPage).String() + "</span>": "<span class=\"current\">" + gvar.New(page.CurrentPage).String() + "</span>",
		"<a":     "<li><a",
		"/a>":    "/a></li>",
		"<span":  "<li class=\"layui-disabled\"><span",
		"/span>": "/span></li>",
	})
	return pageInfo
}
