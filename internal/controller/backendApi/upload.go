package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type cUpload struct{}

var (
	Upload = cUpload{}
)

func (c *cUpload) File(ctx context.Context, req *backendApi.UploadFileReq) (res *backendApi.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	file, err := service.Upload().BackendUploadFile(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})

	if err != nil {
		return nil, err
	}
	res = &backendApi.UploadFileRes{
		Name: file.Name,
		Url:  file.Url,
	}
	return
}
