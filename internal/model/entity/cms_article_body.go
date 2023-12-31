// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsArticleBody is the golang structure for table cms_article_body.
type CmsArticleBody struct {
	Id        uint64      `json:"id"        ` // 自增id
	ArticleId int64       `json:"articleId" ` // 所属文章id
	ChannelId int         `json:"channelId" ` // 所属栏目id
	Body      string      `json:"body"      ` // 文章内容
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
