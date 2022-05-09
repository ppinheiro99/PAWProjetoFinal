package model

import (
	"gorm.io/gorm"
)

type StudentsAndPresentationDone struct {
	gorm.Model
	PresentationDoneID uint   `gorm:"constraint:OnDelete:SET NULL;not null foreignKey:id;references:id"`
	StudentId          uint   ` json:"qr_code"`
	Answers            string `gorm:"type: varchar(255)" json:"answers"`
	Grade              string `gorm:"type: varchar(255)" json:"grade"`
}
