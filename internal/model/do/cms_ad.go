// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsAd is the golang structure of table cms_ad for DAO operations like Where/Data.
type CmsAd struct {
	g.Meta    `orm:"table:cms_ad, do:true"`
	Id        interface{} // 广告id
	ChannelId interface{} // 栏目id
	Name      interface{} // 广告名称
	Link      interface{} // 链接地址
	ImgUrl    interface{} // 图片
	Status    interface{} // 状态(0停用,1显示)
	StartTime *gtime.Time // 广告开始时间
	EndTime   *gtime.Time // 广告结束时间
	Sort      interface{} // 排序
	Remarks   interface{} // 备注
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
