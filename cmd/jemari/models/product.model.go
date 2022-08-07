package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            string         `json:"id" gorm:"primaryKey;"`
	Image         string         `json:"image" gorm:"column:image; type:varchar(255); not null;"`
	Name          string         `json:"name" gorm:"column:name; type:varchar(255); not null;"`
	Address       string         `json:"address" gorm:"column:address; type:varchar(255); not null;"`
	PriceDelivery string         `json:"priceDelivery" gorm:"column:price_delivery; type:varchar(255); not null;"`
	Status        bool           `json:"status" gorm:"column:status; not null;"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeleteAt      gorm.DeletedAt `json:"deleteAt"`
}

func (p *Product) BeforeCreate(db *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
