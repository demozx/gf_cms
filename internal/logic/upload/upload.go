package upload

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sUpload struct{}

var (
	insUpload = sUpload{}
)

func init() {
	service.RegisterUpload(New())
}

func New() *sUpload {
	return &sUpload{}
}

func Upload() *sUpload {
	return &insUpload
}

// BackendSingleUploadFile BackendUploadFile 上传文件
func (*sUpload) SingleUploadFile(ctx context.Context, in model.FileUploadInput, dir string) (out *backendApi.UploadFileRes, err error) {
	serverRoot := service.Util().ServerRoot()
	if err != nil {
		return nil, err
	}
	fullUploadDir := "/upload/" + dir + "/" + gtime.Date()
	fullDir := serverRoot + fullUploadDir
	filename, err := in.File.Save(fullDir, in.RandomName)
	url := fullUploadDir + "/" + filename
	out = &backendApi.UploadFileRes{
		Name: filename,
		Url:  url,
	}
	return
}
