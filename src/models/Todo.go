package models

import (
	"time"
)

type Todo struct {
	ID			int   		`gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"`
	Title    	string 		`gorm:"not null" json:"title"`
	Description	string 		`gorm:"not null" json:"description"`
	Status   	string 		`gorm:"not null" json:"status"`
	CreatedAt 	time.Time 	`gorm:"-" json:"created_at"`
	UpdatedAt 	time.Time	`gorm:"-" json:"updated_at"`
}
