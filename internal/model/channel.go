package model

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
