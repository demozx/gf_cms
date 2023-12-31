// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsArticle is the golang structure for table cms_article.
type CmsArticle struct {
	Id          uint64      `json:"id"          ` // 文章id
	Title       string      `json:"title"       ` // 标题
	ChannelId   int         `json:"channelId"   ` // 所属栏目id
	Keyword     string      `json:"keyword"     ` // 关键词
	Description string      `json:"description" ` // 文章摘要
	Flag        string      `json:"flag"        ` // 属性(p:带图,r:推荐,t:置顶)
	Status      int         `json:"status"      ` // 审核状态(1:已审核,0:未审核)
	Thumb       string      `json:"thumb"       ` // 缩略图
	CopyFrom    string      `json:"copyFrom"    ` // 文章来源
	ClickNum    int         `json:"clickNum"    ` // 点击数
	Sort        int         `json:"sort"        ` // 排序
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 发布时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 编辑时间
	DeletedAt   *gtime.Time `json:"deletedAt"   ` // 删除时间
}
