// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsArticleBodyDao is the data access object for table cms_article_body.
type CmsArticleBodyDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns CmsArticleBodyColumns // columns contains all the column names of Table for convenient usage.
}

// CmsArticleBodyColumns defines and stores column names for table cms_article_body.
type CmsArticleBodyColumns struct {
	Id        string // 自增id
	ArticleId string // 所属文章id
	ChannelId string // 所属栏目id
	Body      string // 文章内容
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// cmsArticleBodyColumns holds the columns for table cms_article_body.
var cmsArticleBodyColumns = CmsArticleBodyColumns{
	Id:        "id",
	ArticleId: "article_id",
	ChannelId: "channel_id",
	Body:      "body",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCmsArticleBodyDao creates and returns a new DAO object for table data access.
func NewCmsArticleBodyDao() *CmsArticleBodyDao {
	return &CmsArticleBodyDao{
		group:   "default",
		table:   "cms_article_body",
		columns: cmsArticleBodyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsArticleBodyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsArticleBodyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsArticleBodyDao) Columns() CmsArticleBodyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsArticleBodyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsArticleBodyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsArticleBodyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
