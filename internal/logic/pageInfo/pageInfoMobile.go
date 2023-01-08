package pageInfo

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// MobilePageInfo mobile分页
func (s *sPageInfo) MobilePageInfo(ctx context.Context, total int, size int) string {
	if total == 0 {
		return ""
	}
	page := g.RequestFromCtx(ctx).GetPage(gconv.Int(total), size)
	currentPage := gconv.String(page.CurrentPage)
	pageInfo := page.GetContent(1)
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<a":  "<li><a",
		"/a>": "/a></li>",
	})
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<span":  "<li><a",
		"/span>": "/a></li>",
	})
	pageInfo = gstr.ReplaceByMap(pageInfo, map[string]string{
		"<a class=\"current\">" + currentPage + "</a>": "<a href='javascript:;'>第" + currentPage + "页</a>",
	})
	return pageInfo
}
