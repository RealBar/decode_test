package fileserver

import (
	"decode_test/pkg/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

func Setup() {
	client, err := oss.New(config.Cfg.FileServerConfig.Endpoint,config.Cfg.FileServerConfig.AccessKey,
		config.Cfg.FileServerConfig.AccessKeySecret,
		oss.SetLogger())
	if err != nil {
		// HandleError(err)
	}
	lsRes, err := client.ListBuckets()
	if err != nil {
		// HandleError(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}
