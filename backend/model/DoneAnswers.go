package model

import (
	"gorm.io/gorm"
)

type DoneAnswers struct {
	gorm.Model
	StudentId  uint   `gorm:"type:int;not null" json:"student_id"`
	Answer     string `gorm:"type:varchar(255);not null" json:"answer"`
	QuestionId uint   `gorm:"type:int;not null" json:"question_id"`
	//DateOfPresentation []time.Time `gorm:"type: date" json:"date"` // Array of dates in which apresentation is held

}
