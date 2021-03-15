package db

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	Title     string
	Text      string
	Forum     bool
	Message   bool
	SectionID int `sql:"type:int REFERENCES sections(id)"`
	OwnerID   int `sql:"type:int REFERENCES users(id)"`
	DirectID  int `sql:"type:int REFERENCES users(id)"`
	MessageID int `sql:"type:int REFERENCES messages(id)"`
}

type MessageFiles struct {
	gorm.Model
	FileID  int `sql:"type:int REFERENCES files(id)"`
	ForumID int `sql:"type:int REFERENCES messages(id)"`
}
