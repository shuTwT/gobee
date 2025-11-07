package model

import "github.com/gofiber/fiber/v2"

type HttpSuccess struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewSuccess(msg string, data any) HttpSuccess {
	return HttpSuccess{
		Code: fiber.StatusOK,
		Msg:  msg,
		Data: data,
	}
}
