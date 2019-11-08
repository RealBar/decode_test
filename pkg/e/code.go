package e

// return code
const (
	OK                     = 0
	ParamInvalid           = 1
	DuplicateMedia         = 101
	BadRequestDelimiter    = 1000
	InternalErrorDelimiter = 9000
	InternalError          = 9999
)

// process exit code
const (
	ExitCodeConfigError  = -1
	ExitCodeConnectError = -2
)
