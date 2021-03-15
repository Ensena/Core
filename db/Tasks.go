package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	Finish      *time.Time
	Start       *time.Time
	SectionID   int `sql:"type:int REFERENCES sections(id)"`
	OwnerID     int `sql:"type:int REFERENCES users(id)"`
	Integrates  int
	Role        int
}

type TaskAnswer struct {
	gorm.Model
	Name         string
	Description  string
	TaskID       int `sql:"type:int REFERENCES tasks(id)"`
	OwnerID      int `sql:"type:int REFERENCES users(id)"`
	Calification int
	Review       string
}

type TaskAnswerUserFriends struct {
	gorm.Model
	TaskAnswerUserID int `sql:"type:int REFERENCES task_answers(id)"`
	OwnerID          int `sql:"type:int REFERENCES users(id)"`
	Calification     int
}

type TaskAnswerUserFiles struct {
	gorm.Model
	FileID int `sql:"type:int REFERENCES files(id)"`
	TaskID int `sql:"type:int REFERENCES task_answers(id)"`
}

type TaskFiles struct {
	gorm.Model
	FileID int `sql:"type:int REFERENCES files(id)"`
	TaskID int `sql:"type:int REFERENCES tasks(id)"`
}
