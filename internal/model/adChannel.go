package model

import "github.com/gogf/gf/v2/util/gmeta"

type AdChannelListItem struct {
	gmeta.Meta  `orm:"table:ad_channel"`
	Id          int    `json:"id" description:"分类id"`
	ChannelName string `json:"channel_name" description:"分类名称"`
	Remarks     string `json:"remarks" description:"备注"`
	Sort        int    `json:"sort" description:"排序"`
}

type AdChannelSortMap struct {
	Id   int `json:"id"`
	Sort int `json:"sort"`
}
