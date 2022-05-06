package model

import (
	"gorm.io/gorm"
)

type Presentations struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null" json:"name"`
	PdfFile []byte `gorm:"type:bytea" json:"pdf_file"`
}
