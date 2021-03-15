package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

//User user from system
type User struct {
	gorm.Model
	Name      string
	LastName  string
	Birthday  *time.Time
	Email     string `gorm:" unique_index"`
	Role      string `gorm:"size:255"` // set field size to 255
	Phone     int
	Address   string
	Domain    string
	Password  string
	AboutMe   string
	Ready     bool
	Public    bool
	Picture   string
	Cover     string
	Rut       string
	MoodleUDP bool
	CanvasID  int
}
