package adList

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/service"
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
