package common

type result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (e result) ErrCode() int {
	return e.Code
}
func (e result) Error() string {
	return e.Msg
}

func NewResult(code int, msg string, data ...interface{}) *result {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	return &result{
		Code: code,
		Msg:  msg,
		Data: d,
	}
}
