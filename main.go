package main

import (
	controller "github.com/ibrahimtrg18/jemari/cmd/jemari/controllers"
	"github.com/ibrahimtrg18/jemari/configs"
	"github.com/labstack/echo"
)

func main() {
	configs.InitEnvironment()
	configs.InitDatabase()
	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(200, "Hello, world!")
	})

	apiGroup := e.Group("/api")

	c := &controller.Controller{DB: configs.DB}

	apiGroup.POST("/user/partner", c.CreatePartnerUserHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
