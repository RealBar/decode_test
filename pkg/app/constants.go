package app

import "time"



const (
	APP_NAME="decode_test"

	MaxFilePathLen = 256
	MaxImageSize   = 2 * 1024 * 1024 //2M
	MaxImageWidth  = 2048
	MaxImageHeight = 2048

	MaxFileSuffixLen = 5
	MaxFileNameLen   = 32

	DBConnectTimeout = 3 * time.Second
	DBRWTimeout      = 1 * time.Second

	DefaultIdcID int64 = 1

	CONFIG_ENV_KEY = "ENV"
	CONFIG_IDC_KEY = "IDC"
	REDIS_PROCESS_KEY = APP_NAME + "_process_id"
)

