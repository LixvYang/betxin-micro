// Topic Error 10xxx
package errmsg

const (
	TOPIC_NOT_FOUND = 10001
)

var TOPIC_ERROR_MSG = map[int]string{
	TOPIC_NOT_FOUND: "topic not found.",
}

func init() {
	for code, msg := range TOPIC_ERROR_MSG {
		codeMsg[code] = msg
	}
}
