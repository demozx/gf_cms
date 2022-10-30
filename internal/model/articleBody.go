package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type ArticleBodyItem struct {
	gmeta.Meta `orm:"table:article_body"`
	Id         uint64      `json:"id"        ` // 自增id
	ArticleId  int64       `json:"articleId" ` // 所属文章id
	ChannelId  int         `json:"channelId" ` // 所属栏目id
	Body       string      `json:"body"      ` // 文章内容
	CreatedAt  *gtime.Time `json:"createdAt" ` //
	UpdatedAt  *gtime.Time `json:"updatedAt" ` //
	DeletedAt  *gtime.Time `json:"deletedAt" ` //
}
