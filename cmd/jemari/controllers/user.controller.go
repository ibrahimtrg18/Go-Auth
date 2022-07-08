package controller

import (
	"net/http"

	model "github.com/ibrahimtrg18/jemari/cmd/jemari/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) CreatePartnerUserHandler(ctx echo.Context) error {
	partnerUser := new(model.PartnerUser)

	if err := ctx.Bind(partnerUser); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	c.DB.Create(partnerUser)

	return ctx.JSON(http.StatusOK, partnerUser)
}
