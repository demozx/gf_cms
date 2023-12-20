// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsRoleAccountDao is the data access object for table cms_role_account.
type CmsRoleAccountDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns CmsRoleAccountColumns // columns contains all the column names of Table for convenient usage.
}

// CmsRoleAccountColumns defines and stores column names for table cms_role_account.
type CmsRoleAccountColumns struct {
	Id        string // ID
	AccountId string // 账户id
	RoleId    string // 角色id
}

// cmsRoleAccountColumns holds the columns for table cms_role_account.
var cmsRoleAccountColumns = CmsRoleAccountColumns{
	Id:        "id",
	AccountId: "account_id",
	RoleId:    "role_id",
}

// NewCmsRoleAccountDao creates and returns a new DAO object for table data access.
func NewCmsRoleAccountDao() *CmsRoleAccountDao {
	return &CmsRoleAccountDao{
		group:   "default",
		table:   "cms_role_account",
		columns: cmsRoleAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsRoleAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsRoleAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsRoleAccountDao) Columns() CmsRoleAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsRoleAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsRoleAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsRoleAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
