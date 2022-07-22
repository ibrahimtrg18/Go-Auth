package main

import (
	controller "github.com/ibrahimtrg18/jemari/cmd/jemari/controllers"
	config "github.com/ibrahimtrg18/jemari/configs"
	util "github.com/ibrahimtrg18/jemari/utils"
	"github.com/labstack/echo"
)

func main() {
	config.InitEnvironment()
	config.InitDatabase()
	e := echo.New()
	v := util.InjectValidate()
	e.Validator = &util.Validator{Validator: v}
	db := config.InjectDatabase()
	c := controller.Controller{DB: db}

	e.GET("/", func(e echo.Context) error {
		return e.String(200, "Hello, world!")
	})

	apiGroup := e.Group("/api")
	apiGroup.POST("/user/partner", c.CreatePartnerUser)
	apiGroup.POST("/user/partner/login", c.LoginPartnerUser)

	e.Logger.Fatal(e.Start(":1323"))
}
