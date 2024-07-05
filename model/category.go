package model

import (
	"time"
)

type Category struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Product   []Product `gorm:"foreignKey:ProductID"`
}
