package db

import (
	"github.com/jinzhu/gorm"
)

//Courses Courses from system
type Courses struct {
	gorm.Model
	Name         string
	Code         string
	InternalCode string
	Enable       bool
	Public       bool
	Scrapped     bool
	// Custom             string
	Description        string
	Video              string
	Picture            string
	OwnerInstitutionID int `sql:"type:int REFERENCES institutions(id)"`
}
