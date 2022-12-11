package friendlyLink

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
)

// PcList pc友情链接
func (s *sFriendlyLink) PcList(ctx context.Context) (out []*entity.CmsFriendlyLink, err error) {
	err = dao.CmsFriendlyLink.Ctx(ctx).
		Where(dao.CmsFriendlyLink.Columns().Status, 1).
		OrderAsc(dao.CmsFriendlyLink.Columns().Sort).
		OrderDesc(dao.CmsFriendlyLink.Columns().Id).Scan(&out)
	if err != nil {
		return nil, err
	}
	return
}
