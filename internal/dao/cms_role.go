// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_cms/internal/dao/internal"
)

// internalCmsRoleDao is internal type for wrapping internal DAO implements.
type internalCmsRoleDao = *internal.CmsRoleDao

// cmsRoleDao is the data access object for table cms_role.
// You can define custom methods on it to extend its functionality as you wish.
type cmsRoleDao struct {
	internalCmsRoleDao
}

var (
	// CmsRole is globally public accessible object for table cms_role operations.
	CmsRole = cmsRoleDao{
		internal.NewCmsRoleDao(),
	}
)

// Fill with you ideas below.
