package fileserver

import (
	"bytes"
	"decode_test/pkg/app"
	"decode_test/pkg/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type FileServer struct {
	cli *oss.Client
}

const (
	BucketNameImage = app.AppName + "_image"
	BucketNameVideo = app.AppName + "_video"
	BucketNameAudio = app.AppName + "_audio"
	BucketNameText  = app.AppName + "_text"
)

func Setup(cfg *config.Config, logger *logrus.Logger) *FileServer {
	client, err := oss.New(cfg.FileServerConfig.Endpoint, cfg.FileServerConfig.AccessKey,
		cfg.FileServerConfig.AccessKeySecret)
	if err != nil {
		logger.WithError(err).Error("init oss client error")
		panic(err)
	}
	lsRes, err := client.ListBuckets()
	if err != nil {
		logger.WithError(err).Error("init oss client error")
		panic(err)
	}
	for _, bucket := range lsRes.Buckets {
		logger.Info("Buckets:", bucket.Name)
	}
	return &FileServer{cli: client}
}

func (f *FileServer) Upload(data []byte, md5 string, t app.MediaType) error {
	bucket, _ := f.cli.Bucket(getBucketFromMediaType(t))
	return bucket.PutObject(md5, bytes.NewReader(data))
}

func getBucketFromMediaType(t app.MediaType) string {
	var res string
	switch t {
	case app.ImageType:
		res = BucketNameImage
	case app.VideoType:
		res = BucketNameVideo
	case app.AudioType:
		res = BucketNameAudio
	case app.TextType:
		res = BucketNameText
	}
	return res
}
