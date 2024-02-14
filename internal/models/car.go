package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	CarModelID uint
	CarModel   CarModel
	Year       int
	Price      float64
	Color      string
	Mileage    int
	VIN        string
	Status     string    // available, sold, etc.
	Features   []Feature `gorm:"many2many:car_features;"`
}
