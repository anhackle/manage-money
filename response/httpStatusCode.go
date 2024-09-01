package response

const (
	ErrCodeSuccess       = 20000
	ErrCodeExternal      = 40000
	ErrCodeParamInvalid  = 40001
	ErrCodeLoginFail     = 40002
	ErrCodeUserHasExists = 40003
	ErrTokenInvalid      = 40100
	ErrCodeInternal      = 50000
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeLoginFail:     "Username or Password invalid",
	ErrCodeInternal:      "Internal server error",
	ErrCodeExternal:      "Bad request",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrTokenInvalid:      "Authorization required",
	ErrCodeUserHasExists: "User existed",
}
