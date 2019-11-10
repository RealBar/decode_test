package e

type BaseError struct {
	userMsg string
	logMsg  string
	code    int
}

func (e BaseError) Error() string {
	return e.logMsg
}

func (e BaseError) UserMsg() string {
	return e.logMsg
}

func (e BaseError) Code() int {
	return e.code
}

func NewError(logMsg string, code int) BaseError {
	return BaseError{logMsg: logMsg, code: code}
}

var (
	ErrAllMediaIDInvalid = NewError("all subMedia ids are invalid", InternalError)
)
