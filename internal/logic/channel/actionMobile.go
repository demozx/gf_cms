package packed

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sChannel) MobileHomeAboutChannel(ctx context.Context, channelId int) (channel *entity.CmsChannel, err error) {
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelId).Scan(&channel)
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, gerror.New("栏目不存在")
	}
	return
}
