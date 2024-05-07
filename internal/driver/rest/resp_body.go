package rest

import "time"

type ResponseBodySuccess struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

type ResponseBodyError struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
	Msg string `json:"msg"`
	Ts  int64  `json:"ts"`
}

// implement later
func NewSuccessResponse(data interface{}) *ResponseBodySuccess {
	return &ResponseBodySuccess{
		Ok:   true,
		Data: data,
		Ts:   time.Now().Unix(),
	}
}

// implement later
func NewErrorResponse(err string, msg string) *ResponseBodyError {
	return &ResponseBodyError{
		Ok:  false,
		Err: err,
		Msg: msg,
		Ts:  time.Now().Unix(),
	}
}
