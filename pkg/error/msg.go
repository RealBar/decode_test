package error

var MsgMap = map[int]string{
	OK:              "OK",
	DUPLICATE_MEDIA: "duplicate media",
	INTERNAL_ERROR:  "server internal error",
	PARAM : "params invalid",
}

func GetErrorMsg(code int) string {
	return MsgMap[code]
}
