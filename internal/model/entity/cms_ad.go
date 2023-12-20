// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsAd is the golang structure for table cms_ad.
type CmsAd struct {
	Id        uint        `json:"id"        ` // 广告id
	ChannelId int         `json:"channelId" ` // 栏目id
	Name      string      `json:"name"      ` // 广告名称
	Link      string      `json:"link"      ` // 链接地址
	ImgUrl    string      `json:"imgUrl"    ` // 图片
	Status    int         `json:"status"    ` // 状态(0停用,1显示)
	StartTime *gtime.Time `json:"startTime" ` // 广告开始时间
	EndTime   *gtime.Time `json:"endTime"   ` // 广告结束时间
	Sort      int         `json:"sort"      ` // 排序
	Remarks   string      `json:"remarks"   ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
