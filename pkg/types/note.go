package types

import "time"

type Note struct {
	Id       string `gorm:"primaryKey;autoIncrement:false;not null"`
	Category string `gorm:"not null"`
	Title    string
	ModTime  time.Time
}
