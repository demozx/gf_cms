// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsArticle is the golang structure of table cms_article for DAO operations like Where/Data.
type CmsArticle struct {
	g.Meta      `orm:"table:cms_article, do:true"`
	Id          interface{} // 文章id
	Title       interface{} // 标题
	ChannelId   interface{} // 所属栏目id
	Keyword     interface{} // 关键词
	Description interface{} // 文章摘要
	Flag        interface{} // 属性(p:带图,r:推荐,t:置顶)
	Status      interface{} // 审核状态(1:已审核,0:未审核)
	Thumb       interface{} // 缩略图
	CopyFrom    interface{} // 文章来源
	ClickNum    interface{} // 点击数
	Sort        interface{} // 排序
	CreatedAt   *gtime.Time // 发布时间
	UpdatedAt   *gtime.Time // 编辑时间
	DeletedAt   *gtime.Time // 删除时间
}
