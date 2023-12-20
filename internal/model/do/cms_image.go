// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsImage is the golang structure of table cms_image for DAO operations like Where/Data.
type CmsImage struct {
	g.Meta      `orm:"table:cms_image, do:true"`
	Id          interface{} // 图片id
	Title       interface{} // 标题
	ChannelId   interface{} // 所属栏目id
	Images      interface{} // 图片们
	Description interface{} // 图片描述
	Flag        interface{} // 属性(r:推荐,t:置顶)
	Status      interface{} // 审核状态(1:启用,0:停用)
	ClickNum    interface{} // 点击数
	Sort        interface{} // 排序
	CreatedAt   *gtime.Time // 发布时间
	UpdatedAt   *gtime.Time // 编辑时间
	DeletedAt   *gtime.Time // 删除时间
}
