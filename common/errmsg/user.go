// user errmsg 20xxx
package errmsg

const (
	USER_NOT_FOUND = 20001
)

var USER_ERROR_MSG = map[int]string{
	USER_NOT_FOUND: "user not found.",
}

func init() {
	for code, msg := range USER_ERROR_MSG {
		codeMsg[code] = msg
	}
}
