package response

const (
	ErrCodeSuccess       = 20000
	ErrCodeInternal      = 20001
	ErrCodeExternal      = 20002
	ErrCodeParamInvalid  = 20003
	ErrCodeLoginFail     = 20004
	ErrTokenInvalid      = 30003
	ErrCodeUserHasExists = 50001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeLoginFail:     "Username or Password invalid",
	ErrCodeInternal:      "Internal server error",
	ErrCodeExternal:      "Bad request",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrTokenInvalid:      "Token is invalid",
	ErrCodeUserHasExists: "User existed",
}
