package model

import "gorm.io/gorm"

// swagger:model
type Subject struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `gorm:"unique;not null" json:"name" `
	Teacher    string `json:"teacher"`
}
