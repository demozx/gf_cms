package adList

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sAdList struct{}
)

var (
	insAdList = sAdList{}
)

func New() *sAdList {
	return &sAdList{}
}

func AdList() *sAdList {
	return &insAdList
}

func init() {
	service.RegisterAdList(New())
}

// Add 添加广告
func (s *sAdList) Add(ctx context.Context, req *backendApi.AdListAddReq) (out interface{}, err error) {
	_, err = dao.CmsAd.Ctx(ctx).Data(req).Insert()
	if err != nil {
		return nil, err
	}
	return
}

// Edit 编辑广告
func (s *sAdList) Edit(ctx context.Context, req *backendApi.AdListEditReq) (out interface{}, err error) {
	update, err := dao.CmsAd.Ctx(ctx).Where(dao.CmsAd.Columns().Id, req.Id).UpdateAndGetAffected(req)
	if err != nil {
		return nil, err
	}
	if update == 0 {
		return nil, gerror.New("编辑的广告不存在")
	}
	return
}

// Delete 删除广告
func (s *sAdList) Delete(ctx context.Context, req *backendApi.AdListDeleteReq) (out interface{}, err error) {
	_, err = dao.CmsAd.Ctx(ctx).WhereIn(dao.CmsAd.Columns().Id, req.Ids).Delete()
	if err != nil {
		return nil, err
	}
	return
}

// GetAdInfoById 根据广告id获取广告信息
func (s *sAdList) GetAdInfoById(ctx context.Context, id int) (out interface{}, err error) {
	var adInfo *entity.CmsAd
	err = dao.CmsAd.Ctx(ctx).Where(dao.CmsAd.Columns().Id, id).Scan(&adInfo)
	if err != nil {
		return nil, err
	}
	if adInfo == nil {
		return nil, gerror.New("广告不存在")
	}
	out = adInfo
	return
}
