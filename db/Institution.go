package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Institution Institution from system
type Institution struct {
	gorm.Model
	Name    string
	Enable  bool
	Public  bool
	Expired *time.Time
	Domain  string
	OwnerID int `sql:"type:int REFERENCES users(id)"`
	AboutMe string
	Picture string
	Cover   string
}

type UsersInstitutions struct {
	gorm.Model
	UserID    int `sql:"type:int REFERENCES users(id)"`
	SectionID int `sql:"type:int REFERENCES institutions(id)"`
	Role      int
}
