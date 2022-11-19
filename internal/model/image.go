package model

import "github.com/gogf/gf/v2/os/gtime"

type ImageGetListInPut struct {
	ChannelId int    `json:"channel_id" description:"栏目ID"`
	StartAt   string `json:"start_at" description:"开始时间"`
	EndAt     string `json:"end_at" description:"结束时间"`
	Keyword   string `json:"keyword" description:"关键词"`
	Page      int    `json:"page" description:"分页码"`
	Size      int    `json:"size" description:"分页数量"`
}

type ImageListItem struct {
	Id            uint64      `json:"id"`           // 图集id
	Title         string      `json:"title"`        // 标题
	ChannelId     int         `json:"channel_id"`   // 所属栏目id
	ChannelName   string      `json:"channel_name"` // 栏目名称
	Description   string      `json:"description"`  // 图集摘要
	Flag          string      `json:"flag"`         // 属性(r:推荐,t:置顶)
	Status        int         `json:"status"`       // 审核状态(1:已审核,0:未审核)
	Thumb         string      `json:"thumb"`        // 主图
	OtherImages   []string    `json:"other_images"` // 其他图
	Images        []string    `json:"images"`       // 图集
	ClickNum      int         `json:"click_num"`    // 点击数
	Sort          int         `json:"sort"`         // 排序
	CreatedAt     *gtime.Time `json:"created_at"`   // 发布时间
	UpdatedAt     *gtime.Time `json:"updated_at"`   // 编辑时间
	FlagR         int         `json:"flag_r" description:"flag推荐"`
	FlagT         int         `json:"flag_t" description:"flag置顶"`
	ImagesStr     string      `json:"images_str"`
	ImagesSrcJson string      `json:"images_src_json"`
}

type ImagesSrcArrItem struct {
	Src   string `json:"src"`
	Title string `json:"title"`
}

type ImageGetListOutPut struct {
	List  []ImageListItem `json:"list" description:"列表"`
	Page  int             `json:"page" description:"分页码"`
	Size  int             `json:"size" description:"分页数量"`
	Total int             `json:"total" description:"数据总数"`
}

type ImageSortMap struct {
	Id   int `json:"id"`
	Sort int `json:"sort"`
}
