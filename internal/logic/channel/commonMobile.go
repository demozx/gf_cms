package packed

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/text/gstr"
)

// MobileListTemplate 获取移动栏目列表模板
func (s *sChannel) MobileListTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error) {
	switch channel.Type {
	case 1:
		// 频道
		switch channel.Model {
		case consts.ChannelModelArticle:
			// 文章
			template = "/mobile/article/list.html"
		case consts.ChannelModelImage:
			// 图集
			template = "/mobile/image/list.html"
		}
	case 2:
		// 单页
		template = "/mobile/single_page/detail.html"
	}
	if len(channel.ListTemplate) > 0 {
		// 后台配置的时候不需要加“/mobile”，程序自动找指定模块下的模板
		template = channel.ListTemplate
		if gstr.HasPrefix(template, "/mobile") {
			template = "/mobile" + template
		}
	}
	return
}

// MobileDetailTemplate 获取移动栏目详情模板
func (s *sChannel) MobileDetailTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error) {
	switch channel.Model {
	case consts.ChannelModelArticle:
		// 文章
		template = "/mobile/article/detail.html"
	case consts.ChannelModelImage:
		// 图集
		template = "/mobile/image/detail.html"
	}
	if len(channel.DetailTemplate) > 0 {
		// 后台配置的时候不需要加“/mobile”，程序自动找指定模块下的模板
		template = channel.DetailTemplate
		if gstr.HasPrefix(template, "/mobile") {
			template = "/mobile" + template
		}
	}
	return
}
