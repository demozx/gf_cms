package model

import "github.com/gogf/gf/v2/util/gmeta"

type AdListItem struct {
	gmeta.Meta  `orm:"table:ad"`
	Id          int    `json:"id" description:"广告id"`
	ChannelId   int    `json:"channel_id" description:"分类id"`
	ChannelName string `json:"channel_name" description:"分类名称"`
	Name        string `json:"name" description:"广告名称"`
	Link        string `json:"link" description:"链接地址"`
	ImgUrl      string `json:"img_url" description:"图片"`
	Status      int    `json:"status" description:"状态(0停用,1显示)"`
	StatusDesc  string `json:"status_desc" description:"状态描述"`
	StartTime   string `json:"start_time" description:"广告开始时间"`
	EndTime     string `json:"end_time" description:"广告结束时间"`
	Sort        int    `json:"sort" description:"排序"`
	Remarks     string `json:"remarks" description:"备注"`
}

type AdBatchStatusItem struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}
