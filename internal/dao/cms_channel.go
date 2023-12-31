// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_cms/internal/dao/internal"
)

// internalCmsChannelDao is internal type for wrapping internal DAO implements.
type internalCmsChannelDao = *internal.CmsChannelDao

// cmsChannelDao is the data access object for table cms_channel.
// You can define custom methods on it to extend its functionality as you wish.
type cmsChannelDao struct {
	internalCmsChannelDao
}

var (
	// CmsChannel is globally public accessible object for table cms_channel operations.
	CmsChannel = cmsChannelDao{
		internal.NewCmsChannelDao(),
	}
)

// Fill with you ideas below.
