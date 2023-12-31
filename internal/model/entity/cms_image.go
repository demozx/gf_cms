// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsImage is the golang structure for table cms_image.
type CmsImage struct {
	Id          uint64      `json:"id"          ` // 图片id
	Title       string      `json:"title"       ` // 标题
	ChannelId   int         `json:"channelId"   ` // 所属栏目id
	Images      string      `json:"images"      ` // 图片们
	Description string      `json:"description" ` // 图片描述
	Flag        string      `json:"flag"        ` // 属性(r:推荐,t:置顶)
	Status      int         `json:"status"      ` // 审核状态(1:启用,0:停用)
	ClickNum    int         `json:"clickNum"    ` // 点击数
	Sort        int         `json:"sort"        ` // 排序
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 发布时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 编辑时间
	DeletedAt   *gtime.Time `json:"deletedAt"   ` // 删除时间
}
