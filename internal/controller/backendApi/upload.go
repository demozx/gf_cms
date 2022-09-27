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
	Upload    = cUpload{}
	typeFile  = "file"
	typeImage = "image"
	typeVideo = "video"
)

func (c *cUpload) SingleFile(ctx context.Context, req *backendApi.UploadFileReq) (res *backendApi.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}
	file, err := service.Upload().SingleUploadFile(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	}, typeFile)
	if err != nil {
		return nil, err
	}
	res = &backendApi.UploadFileRes{
		Name: file.Name,
		Url:  file.Url,
	}
	return
}

func (c *cUpload) SingleImage(ctx context.Context, req *backendApi.UploadFileReq) (res *backendApi.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的图片")
	}
	file, err := service.Upload().SingleUploadFile(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	}, typeImage)
	if err != nil {
		return nil, err
	}
	res = &backendApi.UploadFileRes{
		Name: file.Name,
		Url:  file.Url,
	}
	return
}
func (c *cUpload) SingleVideo(ctx context.Context, req *backendApi.UploadFileReq) (res *backendApi.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的视频")
	}
	file, err := service.Upload().SingleUploadFile(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	}, typeVideo)
	if err != nil {
		return nil, err
	}
	res = &backendApi.UploadFileRes{
		Name: file.Name,
		Url:  file.Url,
	}
	return
}
