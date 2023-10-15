package upload

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/fishtailstudio/imgo"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
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

// SingleUploadFile 上传文件
func (*sUpload) SingleUploadFile(ctx context.Context, in model.FileUploadInput, dir string) (out *backendApi.UploadFileRes, err error) {
	serverRoot := service.Util().ServerRoot()
	os.MkdirAll(serverRoot, 0755)
	os.Chmod(serverRoot, 0755)
	fullUploadDir := "/upload/" + dir + "/" + gtime.Date()
	fullDir := serverRoot + fullUploadDir
	filename, err := in.File.Save(fullDir, in.RandomName)
	if err != nil {
		return nil, err
	}
	url := fullUploadDir + "/" + filename
	// 图片质量压缩
	imageQuality := gconv.Int(service.Util().GetSetting("image_quality"))
	if imageQuality >= 1 && imageQuality <= 100 {
		imgo.Load(serverRoot+url).Save(serverRoot+url, imageQuality)
	}
	out = &backendApi.UploadFileRes{
		Name: filename,
		Url:  url,
	}
	return
}
