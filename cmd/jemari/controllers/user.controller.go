package controller

import (
	"net/http"

	"github.com/ibrahimtrg18/jemari/cmd/jemari/dto"
	model "github.com/ibrahimtrg18/jemari/cmd/jemari/models"
	util "github.com/ibrahimtrg18/jemari/utils"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (c *Controller) CreatePartnerUser(ctx echo.Context) error {
	registerPartnerUser := new(dto.RegisterPartnerUser)
	partnerUser := &model.PartnerUser{}

	if err := ctx.Bind(registerPartnerUser); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err := ctx.Validate(registerPartnerUser); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	results := c.DB.Where("pic_email = ?", registerPartnerUser.Email).Find(&partnerUser)

	if results.RowsAffected > 0 {
		u := util.Response{Code: http.StatusConflict, Message: "Failed Create Partner Account!"}
		return u.ResponseFailed(ctx)
	}

	partnerUser = &model.PartnerUser{
		FirstName: registerPartnerUser.FirstName,
		LastName:  registerPartnerUser.LastName,
		Email:     registerPartnerUser.Email,
		Password:  registerPartnerUser.Password,
		Name:      registerPartnerUser.Name,
		Type:      registerPartnerUser.Type,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(partnerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	partnerUser.Password = string(hashedPassword)

	c.DB.Create(partnerUser)

	u := util.Response{Code: http.StatusCreated, Message: "Successfully Create Partner Account!", Data: partnerUser}

	return u.ResponseSuccess(ctx)
}

func (c *Controller) LoginPartnerUser(ctx echo.Context) error {
	partnerUserDTO := new(model.PartnerUser)

	if err := ctx.Bind(partnerUserDTO); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	partnerUser := new(model.PartnerUser)

	if err := c.DB.Where("pic_email = ?", partnerUserDTO.Email).First(&partnerUser).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Not found!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(partnerUser.Password), []byte(partnerUserDTO.Password)); err != nil {
		return ctx.String(http.StatusInternalServerError, "Wrong password!")
	}

	payloadToken := map[string]interface{}{
		"id": partnerUser.ID,
	}

	accessToken, err := util.CreateToken(payloadToken, 1)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	u := util.Response{Code: http.StatusOK, Message: "Successfully Login Account!", Data: map[string]interface{}{
		"accessToken": accessToken,
	}}

	return u.ResponseSuccess(ctx)
}
