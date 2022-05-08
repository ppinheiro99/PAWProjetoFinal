package model

import (
	"gorm.io/gorm"
)

type SubjectPresentations struct {
	gorm.Model
	PresentationID uint `gorm:"constraint:OnDelete:SET NULL; not null foreignKey:id;references:id"`
	SubjectID      uint `gorm:"constraint:OnDelete:SET NULL; foreignKey:id;references:id not null" `
}
