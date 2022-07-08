package service

import (
	"fmt"

	model "github.com/ibrahimtrg18/jemari/cmd/jemari/models"
)

type UserService interface {
	CreateUser(partnerUser *model.PartnerUser)
}

func CreateUser(partnerUser *model.PartnerUser) {
	fmt.Println(partnerUser)
}
