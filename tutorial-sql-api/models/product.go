package models

import (
	"time"
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type Product struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null;unique;type:varchar(191)"`
	Brand string `gorm:"not null;unique;type:varchar(191)"`
	UserID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product before create() // this line in product.go")

	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}

	return
}