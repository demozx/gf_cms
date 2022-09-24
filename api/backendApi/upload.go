package backendApi

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `tags:"BackendApi" method:"post" mime:"multipart/form-data" summary:"后台上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}
type UploadFileRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"文件url"`
}
