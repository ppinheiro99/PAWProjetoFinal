package model

import (
	"gorm.io/gorm"
)

type DoneAnswers struct {
	gorm.Model
	StudentUsername string `gorm:"type:varchar(255);not null" json:"student_username"`
	Answer          string `gorm:"type:varchar(255);not null" json:"answer"`
	QuestionId      uint   `gorm:"type:int;not null" json:"question_id"`
	Was_Right       bool   `gorm:"type:bool;not null" json:"was_right"`
	//DateOfPresentation []time.Time `gorm:"type: date" json:"date"` // Array of dates in which apresentation is held

}
