package util

import (
	"github.com/labstack/echo"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) ResponseSuccess(ctx echo.Context) error {
	if r.Code > 0 {
		return ctx.JSON(r.Code, r)
	}

	r.Code = 200
	return ctx.JSON(r.Code, r)
}

func (r Response) ResponseFailed(ctx echo.Context) error {
	if r.Code > 0 {
		return ctx.JSON(r.Code, r)
	}

	r.Code = 400
	return ctx.JSON(r.Code, r)
}
