package utils

import "time"

const (
	MaxFilePathLen = 256
	MaxImageSize   = 2 * 1024 * 1024
	MaxImageWidth  = 2048
	MaxImageHeight = 2048

	MaxFileSuffixLen = 5
	MaxFileNameLen = 32

	DBConnectTimeout = 3 * time.Second
)
