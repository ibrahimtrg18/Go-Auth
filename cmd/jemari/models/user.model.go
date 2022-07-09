package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PartnerUser struct {
	Pic
	Venture
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeleteAt  gorm.DeletedAt `json:"deleteAt"`
}

type Pic struct {
	ID              string `json:"id" gorm:"primaryKey;"`
	FirstName       string `json:"firstName" gorm:"column:pic_first_name; type:varchar(255); not null;"`
	LastName        string `json:"lastName" gorm:"column:pic_last_name; type:varchar(255); not null;"`
	Email           string `json:"email" gorm:"column:pic_email; type:varchar(255); not null;"`
	Password        string `json:"password" gorm:"column:password; not null;"`
	ConfirmPassword string `json:"confirmPassword"`
}

type Venture struct {
	Name string `json:"name" gorm:"column:venture_name; type:varchar(255); not null;"`
	Type string `json:"type" gorm:"column:venture_type; type:varchar(255); not null;"`
}

func (user *PartnerUser) BeforeCreate(db *gorm.DB) error {
	user.ID = uuid.New().String()
	return nil
}
