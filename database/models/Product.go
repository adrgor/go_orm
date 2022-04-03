package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string
	Price float32
	Availbe bool
	Category string
}