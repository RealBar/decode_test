package e

var MsgMap = map[int]string{
	OK:             "OK",
	DuplicateMedia: "duplicate media",
	InternalError:  "server internal e",
	ParamInvalid:   "params invalid",
}

func GetErrorMsg(code int) string {
	return MsgMap[code]
}
