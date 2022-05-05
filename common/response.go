package common

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (e response) ErrCode() int {
	return e.Code
}
func (e response) Error() string {
	return e.Msg
}

func NewResponse(code int, msg string, data ...interface{}) *response {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	return &response{
		Code: code,
		Msg:  msg,
		Data: d,
	}
}
