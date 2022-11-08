package backendApi

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type AdListIndexReq struct {
	g.Meta    `tags:"Backend" method:"post" summary:"后台广告列表"`
	ChannelId int `json:"channel_id" description:"分类id"`
	model.PageSizeReq
}
type AdListIndexRes struct {
	List  []*model.AdListItem `json:"list" description:"后台广告列表接口结果"`
	Page  int                 `json:"page" description:"分页码"`
	Size  int                 `json:"size" description:"分页数量"`
	Total int                 `json:"total" description:"数据总数"`
}

type AdListAddReq struct {
	g.Meta    `tags:"Backend" method:"post" summary:"后台广告添加"`
	ChannelId int    `json:"channel_id" dc:"栏目id" arg:"true" v:"required|min:1#请选择所属分类|分类ID不能为0"`
	Name      string `json:"name" dc:"广告名称" arg:"true" v:"required#请输入广告名称"`
	Link      string `json:"link" dc:"链接地址" arg:"true" v:"required|url#请输入链接地址|链接地址格式错误"`
	ImgUrl    string `json:"img_url" dc:"图片" v:""`
	Status    int    `json:"status" dc:"状态(0停用,1显示)" v:"required|in:0,1#请选择是否启用|启用值不合法" d:"0"`
	StartTime string `json:"start_time" dc:"广告开始时间" v:"required|datetime#请输入广告开始时间|广告开始时间格式错误"`
	EndTime   string `json:"end_time" dc:"广告结束时间" v:"required|datetime#请输入广告结束时间|广告结束时间格式错误"`
	Remarks   string `json:"remarks" dc:"备注" v:""`
}
type AdListAddRes struct{}
