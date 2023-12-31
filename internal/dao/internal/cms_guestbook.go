// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsGuestbookDao is the data access object for table cms_guestbook.
type CmsGuestbookDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns CmsGuestbookColumns // columns contains all the column names of Table for convenient usage.
}

// CmsGuestbookColumns defines and stores column names for table cms_guestbook.
type CmsGuestbookColumns struct {
	Id        string //
	Name      string // 留言者姓名
	Tel       string // 留言者电话
	Content   string // 留言内容
	From      string // 来源：1、电脑端，2、移动端
	Ip        string // 留言者ip
	Address   string // 留言者归属地
	Status    string // 留言状态：0、未读，1、已读
	CreatedAt string //
	UpdatedAt string //
}

// cmsGuestbookColumns holds the columns for table cms_guestbook.
var cmsGuestbookColumns = CmsGuestbookColumns{
	Id:        "id",
	Name:      "name",
	Tel:       "tel",
	Content:   "content",
	From:      "from",
	Ip:        "ip",
	Address:   "address",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCmsGuestbookDao creates and returns a new DAO object for table data access.
func NewCmsGuestbookDao() *CmsGuestbookDao {
	return &CmsGuestbookDao{
		group:   "default",
		table:   "cms_guestbook",
		columns: cmsGuestbookColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsGuestbookDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsGuestbookDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsGuestbookDao) Columns() CmsGuestbookColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsGuestbookDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsGuestbookDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsGuestbookDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
