package db

import "github.com/jinzhu/gorm"

//Files Files upload from system
type Files struct {
	gorm.Model
	Name    string
	URL     string
	Enable  string
	OwnerID int `sql:"type:int REFERENCES users(id)"`
}

type FilesSections struct {
	gorm.Model
	FileID    int `sql:"type:int REFERENCES files(id)"`
	SectionID int `sql:"type:int REFERENCES sections(id)"`
	Role      int
}
