// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsAdChannel is the golang structure for table cms_ad_channel.
type CmsAdChannel struct {
	Id          uint        `json:"id"          ` //
	ChannelName string      `json:"channelName" ` //
	Remarks     string      `json:"remarks"     ` //
	Sort        int         `json:"sort"        ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
