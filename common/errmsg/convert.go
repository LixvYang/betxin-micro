package errmsg

func GetErrMsg(code int) string {
	return codeMsg[code]
}
