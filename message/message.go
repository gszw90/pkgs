package message

const (
	CodeSuccess  = 200
	CodeFail     = 400
	CodeUnauthorized = 401
	CodeNotFound = 404

	MsgSuccess  = "success"
	MsgFail     = "fail"
	MsgUnauthorized = "unauthorized"
	MsgNotFound = "not found"
)

// Message 格式化返回消息
type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewMessage return message
func NewMessage(code int, msg string, data interface{}) Message {
	return Message{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// SuccessMessage success
func SuccessMessage(data interface{}, msg string) Message {
	if msg == "" {
		msg = MsgSuccess
	}
	return NewMessage(CodeSuccess, msg, data)
}

// FailMessage fail
func FailMessage(msg string, data interface{}) Message {
	if msg == "" {
		msg = MsgFail
	}
	return NewMessage(CodeFail, msg, data)
}

// NotFoundMessage not found
func NotFoundMessage(msg string, data interface{}) Message {
	if msg == "" {
		msg = MsgNotFound
	}
	return NewMessage(CodeNotFound, msg, data)
}

func UnauthorizedMessage() Message {
	return NewMessage(CodeUnauthorized, MsgUnauthorized, nil)
}
