// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsArticleBody is the golang structure of table cms_article_body for DAO operations like Where/Data.
type CmsArticleBody struct {
	g.Meta    `orm:"table:cms_article_body, do:true"`
	Id        interface{} // 自增id
	ArticleId interface{} // 所属文章id
	ChannelId interface{} // 所属栏目id
	Body      interface{} // 文章内容
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
