package model

import (
	"gorm.io/gorm"
)

type PresentationAndQuestion struct {
	gorm.Model
	PresentationID uint `gorm:"not null foreignKey:id;references:id"`
	QuestionID     uint `gorm:"foreignKey:id;references:id not null" json:"question"`
}
