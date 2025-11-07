package model

type HttpError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg string) HttpError {
	return HttpError{
		Code: code,
		Msg:  msg,
	}
}
