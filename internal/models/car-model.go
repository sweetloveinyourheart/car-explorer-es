package models

import "gorm.io/gorm"

type CarModel struct {
	gorm.Model
	ProducerID uint
	Producer   Producer
	Name       string
}
