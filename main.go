package main

import (
	controller "github.com/ibrahimtrg18/jemari/cmd/jemari/controllers"
	config "github.com/ibrahimtrg18/jemari/configs"
	"github.com/labstack/echo"
)

func main() {
	config.InitEnvironment()
	config.InitDatabase()
	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(200, "Hello, world!")
	})

	apiGroup := e.Group("/api")

	db := config.InjectDatabase()

	c := controller.Controller{DB: db}

	apiGroup.POST("/user/partner", c.CreatePartnerUser)
	apiGroup.POST("/user/partner/login", c.LoginPartnerUser)

	e.Logger.Fatal(e.Start(":1323"))
}
