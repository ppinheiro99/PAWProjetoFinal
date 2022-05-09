package model

import (
	"gorm.io/gorm"
)

type PresentationDone struct {
	gorm.Model
	PresentationID     uint   `gorm:"constraint:OnDelete:SET NULL;not null foreignKey:id;references:id"`
	QRCode             []byte `gorm:"type:bytea" json:"qr_code"`
	DateOfPresentation string `gorm:"type: date" json:"date"`
}
