package model

type GuestbookGetListOutputItem struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	Content   string `json:"content"`
	From      int    `json:"from"`
	FromDesc  string `json:"from_desc"`
	Ip        string `json:"ip"`
	Address   string `json:"address"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GuestbookGetListOutput struct {
	List  []*GuestbookGetListOutputItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type BaiduAddressByIp struct {
	Address string `json:"address"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Content BaiduAddressByIpContent
}

type BaiduAddressByIpContent struct {
	Address string `json:"address"`
}
