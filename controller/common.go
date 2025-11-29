package controller

import "github.com/kangyueyue/go-ai/common/code"

// Response resp
type Response struct {
	StatusCode code.Code `json:"status_code"`
	StatusMsg  string    `json:"status_msg"`
}

// CodeOf code
func (r *Response) CodeOf(code code.Code) Response {
	if nil == r {
		r = new(Response)
	}
	r.StatusCode = code
	r.StatusMsg = code.Msg()
	return *r
}

// Success 成功
func (r *Response) Success() {
	r.CodeOf(code.CodeSuccess)
}
