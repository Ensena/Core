package db

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//Auth Auth from system
type Auth struct {
	gorm.Model
	System   int
	User     string
	Password string
	Custom   postgres.Jsonb
	UserID   int `sql:"type:int REFERENCES users(id)"`
}
