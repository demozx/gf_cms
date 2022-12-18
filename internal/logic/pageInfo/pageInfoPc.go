package pageInfo

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcPageInfo pc分页
func (s *sPageInfo) PcPageInfo(ctx context.Context, total int64, size int) string {
	if total == 0 {
		return ""
	}
	page := g.RequestFromCtx(ctx).GetPage(gconv.Int(total), size)
	currentPage := gconv.String(page.CurrentPage)
	pageInfo := page.GetContent(3)
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<ul>": "<ul class=\"pagination\">",
	})
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<a":  "<li><a",
		"/a>": "/a></li>",
	})
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<span":  "<li><a",
		"/span>": "/a></li>",
	})
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"class=\"GPageSpan\"": "",
	})
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<a >" + currentPage + "</a>": "<li class=\"active\"><a>" + currentPage + "</a><li>",
	})
	return pageInfo
}
