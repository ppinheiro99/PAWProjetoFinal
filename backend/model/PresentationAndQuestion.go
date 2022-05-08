package model

import (
	"gorm.io/gorm"
)

type PresentationAndQuestion struct {
	gorm.Model
	PresentationID uint `gorm:"constraint:OnDelete:SET NULL;not null foreignKey:id;references:id"`
	QuestionID     uint `gorm:"constraint:OnDelete:SET NULL;foreignKey:id;references:id not null" json:"question"`
}
