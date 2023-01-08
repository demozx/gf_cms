package packed

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/text/gstr"
)

// PcListTemplate 获取pc栏目列表模板
func (s *sChannel) PcListTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error) {
	switch channel.Type {
	case 1:
		// 频道
		switch channel.Model {
		case consts.ChannelModelArticle:
			// 文章
			template = "/pc/article/list.html"
		case consts.ChannelModelImage:
			// 图集
			template = "/pc/image/list.html"
		}
	case 2:
		// 单页
		template = "/pc/single_page/detail.html"
	}
	if len(channel.ListTemplate) > 0 {
		// 后台配置的时候需要加“{module}”，程序自动找指定模块下的模板
		template = channel.ListTemplate
		if gstr.HasPrefix(template, "{module}") {
			template = gstr.Replace(template, "{module}", "/pc")
		}
	}
	return
}

// PcDetailTemplate 获取pc栏目详情模板
func (s *sChannel) PcDetailTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error) {
	switch channel.Model {
	case consts.ChannelModelArticle:
		// 文章
		template = "/pc/article/detail.html"
	case consts.ChannelModelImage:
		// 图集
		template = "/pc/image/detail.html"
	}
	if len(channel.DetailTemplate) > 0 {
		// 后台配置的时候需要加“{module}”，程序自动找指定模块下的模板
		template = channel.DetailTemplate
		if gstr.HasPrefix(template, "{module}") {
			template = gstr.Replace(template, "{module}", "/pc")
		}
	}
	return
}
