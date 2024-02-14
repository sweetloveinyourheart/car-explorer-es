package models

import "gorm.io/gorm"

type Producer struct {
	gorm.Model
	Name           string
	Country        string
	FoundationYear int
}
