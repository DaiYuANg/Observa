package model

import "github.com/gofiber/fiber/v2"

type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Resp[T any] struct {
	Body R[T] `json:"body"`
}

// 通用成功
func OK[T any](c *fiber.Ctx, data T) error {
	return c.JSON(R[T]{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func ReturnOk[A R[T], T any](data T) *Resp[T] {
	return &Resp[T]{
		Body: R[T](A{
			Code:    0,
			Message: "success",
			Data:    data,
		}),
	}
}

// 自定义成功消息
func OKMsg[T any](c *fiber.Ctx, msg string, data T) error {
	return c.JSON(R[T]{
		Code:    0,
		Message: msg,
		Data:    data,
	})
}

// 通用失败（带状态码）
func Fail(c *fiber.Ctx, code int, msg string) error {
	return c.JSON(R[any]{
		Code:    code,
		Message: msg,
	})
}

// 快速 400 错误
func BadRequest(c *fiber.Ctx, msg string) error {
	return Fail(c, 400, msg)
}

// 快速 500 错误
func ServerError(c *fiber.Ctx, msg string) error {
	return Fail(c, 500, msg)
}
