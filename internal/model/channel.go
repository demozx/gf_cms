package model

import "gf_cms/internal/model/entity"

type ChannelNavigationListItem struct {
	entity.CmsChannel
	HasChildren   bool                         `json:"has_children"`   // 是否有子集
	Children      []*ChannelNavigationListItem `json:"children"`       // 子集
	ChannelRouter string                       `json:"channel_router"` // 处理后的可以直接使用的url地址
	TriggerType   string                       `json:"trigger_type"`   // 处理后的可以直接使用的链接打开方式（当前标签，新标签）
	Current       bool                         `json:"current"`        // 是不是当前栏目id
	Highlight     bool                         `json:"highlight"`      // 是不是高亮栏目
}

// ChannelBackendApiListItem 后台栏目分类接口列表数据
type ChannelBackendApiListItem struct {
	Id        int                          `json:"id"`
	Pid       int                          `json:"pid"`
	Level     int                          `json:"level"`
	Name      string                       `json:"name"`
	Status    int                          `json:"status"`
	Selected  uint8                        `json:"selected"`
	Model     string                       `json:"model"`
	ModelDesc string                       `json:"modelDesc"`
	Children  []*ChannelBackendApiListItem `json:"children"`
}

type ChannelBackendApiListItemChildren struct {
	Id     int    `json:"id"`
	Pid    int    `json:"pid"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type ChannelTreeItem struct {
	Id    int    `json:"id"`
	Pid   int    `json:"pid"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type ChannelTDK struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
}

type ChannelCrumbs struct {
	Name   string `json:"name"`
	Router string `json:"router"`
}
