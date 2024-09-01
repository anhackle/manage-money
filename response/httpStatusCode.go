package response

const (
	ErrCodeSuccess       = 20000
	ErrCodeInternal      = 50000
	ErrCodeExternal      = 40000
	ErrCodeParamInvalid  = 20003
	ErrCodeLoginFail     = 20004
	ErrTokenInvalid      = 40100
	ErrCodeUserHasExists = 50001
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
