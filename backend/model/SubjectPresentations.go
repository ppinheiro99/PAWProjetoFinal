package model

import (
	"gorm.io/gorm"
)

type SubjectPresentations struct {
	gorm.Model
	PresentationID uint `gorm:"not null foreignKey:id;references:id"`
	SubjectID      uint `gorm:"foreignKey:id;references:id not null" `
}
