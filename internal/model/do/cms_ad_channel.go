// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsAdChannel is the golang structure of table cms_ad_channel for DAO operations like Where/Data.
type CmsAdChannel struct {
	g.Meta      `orm:"table:cms_ad_channel, do:true"`
	Id          interface{} //
	ChannelName interface{} //
	Remarks     interface{} //
	Sort        interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
