package model

import (
	"gorm.io/gorm"
)

type Questions struct {
	gorm.Model
	Question      string `gorm:"type:varchar(255);not null" json:"question"`
	Answers       string `gorm:"type:varchar(255);not null" json:"answers"`
	CorrectAnswer string `gorm:"type:varchar(255);not null" json:"correct_answer"`
}
