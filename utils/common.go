package util

import "github.com/labstack/echo"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) ResponseSuccess(ctx echo.Context) error {
	return ctx.JSON(r.Code, r)
}
