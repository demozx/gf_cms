// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_cms/internal/service/internal/dao/internal"
)

// internalCmsSystemConfigCatDao is internal type for wrapping internal DAO implements.
type internalCmsSystemConfigCatDao = *internal.CmsSystemConfigCatDao

// cmsSystemConfigCatDao is the data access object for table cms_system_config_cat.
// You can define custom methods on it to extend its functionality as you wish.
type cmsSystemConfigCatDao struct {
	internalCmsSystemConfigCatDao
}

var (
	// CmsSystemConfigCat is globally public accessible object for table cms_system_config_cat operations.
	CmsSystemConfigCat = cmsSystemConfigCatDao{
		internal.NewCmsSystemConfigCatDao(),
	}
)

// Fill with you ideas below.
