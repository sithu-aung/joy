package models

import (
	"gorm.io/gorm"
)

// Quiz represents a quiz for a lesson
type Quiz struct {
	gorm.Model
	Title     string `gorm:"not null"`
	LessonID  uint
	Questions []Question
}

// Question represents a question in a quiz
type Question struct {
	gorm.Model
	QuizID      uint
	Text        string `gorm:"not null"`
	Options     []Option
	CorrectOption uint
	Explanation string
}

// Option represents an answer option for a question
type Option struct {
	gorm.Model
	QuestionID uint
	Text       string `gorm:"not null"`
	Order      int
}