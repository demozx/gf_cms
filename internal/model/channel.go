package model

// ChannelBackendApiListItem 后台栏目分类接口列表数据
type ChannelBackendApiListItem struct {
	Id       int                                  `json:"id"`
	Pid      int                                  `json:"pid"`
	Name     string                               `json:"name"`
	Status   int                                  `json:"status"`
	Children []*ChannelBackendApiListItemChildren `json:"children"`
}

type ChannelBackendApiListItemChildren struct {
	Id     int    `json:"id"`
	Pid    int    `json:"pid"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
