package backendApi

import "github.com/gogf/gf/v2/frame/g"

type ChannelIndexApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目分类列表"`
}
type ChannelIndexApiRes struct {
}

type ChannelStatusApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目启用禁用"`
	Id     int `name:"id" brief:"栏目ID" dc:"栏目ID"  arg:"true" v:"required#请选择栏目ID"`
}
type ChannelStatusApiRes struct {
}

type ChannelDeleteApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目删除"`
	Id     int `name:"id" brief:"栏目ID" dc:"栏目ID"  arg:"true" v:"required#请选择栏目ID"`
}
type ChannelDeleteApiRes struct {
}

type ChannelAddApiReq struct {
	g.Meta         `tags:"Backend" method:"post" summary:"栏目添加、编辑"`
	Pid            int    `name:"pid" brief:"父级分类ID" dc:"父级分类ID" arg:"true" d:"0" v:"required#请选择父级分类ID"`
	Name           string `name:"name" brief:"分类名称" dc:"分类名称"  arg:"true" v:"required#分类名称"`
	Thumb          string `name:"thumb" brief:"栏目缩略图" dc:"栏目缩略图"  arg:"true" v:""`
	Sort           int    `name:"sort" brief:"排序" dc:"排序" arg:"true" d:"100" v:"required#请输入排序"`
	Status         int    `name:"status" brief:"状态" dc:"状态" arg:"true" d:"0" v:"required|in:0,1#请填写状态|状态不合法"`
	Description    string `name:"description" brief:"栏目简介" dc:"栏目简介"  arg:"true" v:""`
	Type           int    `name:"type" brief:"栏目类型" dc:"栏目类型" arg:"true" d:"0" v:"required|in:1,2,3#请选择栏目类型|栏目类型不合法"`
	LinkUrl        string `name:"link_url" brief:"链接地址" dc:"链接地址"  arg:"true" v:""`
	LinkTrigger    int    `name:"link_trigger" brief:"打开方式" dc:"打开方式" arg:"true" d:"0" v:"required|in:0,1#请选择打开方式|打开方式不合法"`
	Model          string `name:"model" brief:"模型" dc:"模型"  arg:"true" v:"required#请选择模型"`
	ListRouter     string `name:"list_router" brief:"列表路由" dc:"列表路由"  arg:"true" v:""`
	DetailRouter   string `name:"detail_router" brief:"详情路由" dc:"详情路由"  arg:"true" v:""`
	ListTemplate   string `name:"list_template" brief:"列表模板" dc:"列表模板"  arg:"true" v:""`
	DetailTemplate string `name:"detail_template" brief:"详情模板" dc:"详情模板"  arg:"true" v:""`
}
type ChannelAddApiRes struct {
}

type ChannelEditApiReq struct {
	Id int `name:"id" brief:"分类ID" dc:"分类ID" arg:"true" d:"0" v:"required|min:1#请选择父级分类ID|分类ID不能为0"`
	ChannelAddApiReq
}
type ChannelEditApiRes struct {
}
