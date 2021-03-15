package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Section struct {
	gorm.Model
	Code         string
	InternalCode string
	Semester     int
	Year         int
	Section      int
	Start        *time.Time
	Finish       *time.Time
	Enable       bool
	Public       bool
	Scrapped     bool
	// Custom       string
	OwnerID  int `sql:"type:int REFERENCES users(id)"`
	CourseID int `sql:"type:int REFERENCES courses(id)"`
}

type UsersSections struct {
	gorm.Model
	UserID    int `sql:"type:int REFERENCES users(id)"`
	SectionID int `sql:"type:int REFERENCES sections(id)"`
	Role      int
}
