package adList

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

// PcHomeListByChannelId pc首页banner
func (s *sAdList) PcHomeListByChannelId(ctx context.Context, channelId int) (out []*entity.CmsAd, err error) {
	var adList []*entity.CmsAd
	m := dao.CmsAd.Ctx(ctx)
	err = m.Where(dao.CmsAd.Columns().Status, 1).
		Where(dao.CmsAd.Columns().ChannelId, channelId).
		OrderAsc(dao.CmsAd.Columns().Sort).
		OrderAsc(dao.CmsAd.Columns().Id).
		Where(
			m.Builder().Where(
				m.Builder().WhereLTE(dao.CmsAd.Columns().StartTime, gtime.Now()).WhereGTE(dao.CmsAd.Columns().EndTime, gtime.Now()),
			).WhereOr("`start_time` = `end_time`"),
		).Scan(&adList)
	if err != nil {
		return nil, err
	}
	for index, item := range adList {
		if item.Link == "" {
			adList[index].Link = "javascript:;"
		}
	}
	out = adList
	return
}
