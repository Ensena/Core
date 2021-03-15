package db

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//Auth Auth from system
type App struct {
	gorm.Model
	Name        string
	Description string
	Key         string
	InternalKey string
	Custom      postgres.Jsonb
	OwnerID     int `sql:"type:int REFERENCES users(id)"`
	Admin       postgres.Jsonb
	URL         string
	Production  bool
	Telegram    int
}

type UsersApps struct {
	gorm.Model
	UserID int `sql:"type:int REFERENCES users(id)"`
	AppID  int `sql:"type:int REFERENCES apps(id)"`
	Enable bool
	Data   postgres.Jsonb
}
