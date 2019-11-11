package app

import "time"

type MediaStatus uint8
type MediaType uint8
type MediaFormat uint8
type Env string

const (
	AppName = "decode_test"

	MaxFilePathLen = 256
	MaxImageSize   = 2 * 1024 * 1024 //2M
	MaxImageWidth  = 4096
	MaxImageHeight = 4096
	ThumbnailThreshold = 200

	MaxFileSuffixLen = 5
	MaxFileNameLen   = 32

	DBConnectTimeout = 3 * time.Second
	DBRWTimeout      = 1 * time.Second

	ConfigEnvKey    = "ENV"
	ConfigIdcKey    = "IDC"
	RedisProcessKey = AppName + "_process_id"
	CtxKeyOwnerID = "owner_od"
	CtxKeyRequestID = "x-request-id"
	CtxKeyResponseID = "x-response-id"
	CtxKeyLogger = "logger"

	ENV_DEV  Env = "DEV"
	ENV_TEST Env = "TEST"
	ENV_LIVE Env = "LIVE"

	DefaultListenPort            = 8001
	DefaultIdcID           int64 = 1
	DefaultEnv                   = ENV_DEV
	DefaultMaxHeaderBytes        = 1 << 20
	DefaultReadTimeout        = 10 * time.Second
	DefaultWriteTimeout       = 10 * time.Second
)

const (
	_ = MediaStatus(iota)
	MediaStatusNormal
	MediaStatusDeleted
)
const (
	_ = MediaType(iota)
	ImageType
	VideoType
	AudioType
	TextType
)

const (
	_ = MediaFormat(iota)
	Jpeg
	Gif
	Png
	Mp4
	Mp3
	SimpleText
	Unknown MediaFormat = 999
)


