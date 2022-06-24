package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(200, "Hello, world!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
