// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsGuestbook is the golang structure for table cms_guestbook.
type CmsGuestbook struct {
	Id        uint        `json:"id"        ` //
	Name      string      `json:"name"      ` // 留言者姓名
	Tel       string      `json:"tel"       ` // 留言者电话
	Content   string      `json:"content"   ` // 留言内容
	From      int         `json:"from"      ` // 来源：1、电脑端，2、移动端
	Ip        string      `json:"ip"        ` // 留言者ip
	Address   string      `json:"address"   ` // 留言者归属地
	Status    int         `json:"status"    ` // 留言状态：0、未读，1、已读
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
