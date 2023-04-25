// category errmsg
package errmsg

const (
	CATEGORY_NOT_FOUND = 30001
)

var CATEGORY_ERROR_MSG = map[int]string{
	CATEGORY_NOT_FOUND: "category not found.",
}

func init() {
	for code, msg := range CATEGORY_ERROR_MSG {
		codeMsg[code] = msg
	}
}
