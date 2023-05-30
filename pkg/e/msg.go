package e

var MsgFlags = map[uint]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求的参数错误",
}

func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	} else {
		return MsgFlags[Error]
	}
}
