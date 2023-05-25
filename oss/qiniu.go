package oss

import (
	"context"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var qiniuConfig *ConfigQiniu

type ConfigQiniu struct {
	Ak     string
	Sk     string
	Bucket string
	Domain string
}

func SetQiniuConfig(c *ConfigQiniu) {
	qiniuConfig = c
}

var QiniuOss = &qiniuOss{}

type qiniuOss struct{}

func (o qiniuOss) Upload(key, file string) (url string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: qiniuConfig.Bucket,
	}
	mac := qbox.NewMac(qiniuConfig.Ak, qiniuConfig.Sk)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, file, &putExtra)
	if err != nil {
		return
	}
	url = qiniuConfig.Domain + key
	return
}
