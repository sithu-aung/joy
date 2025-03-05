package models

import (
	"gorm.io/gorm"
)

// UserProgress tracks a user's progress through courses and lessons
type UserProgress struct {
	gorm.Model
	UserID       uint
	CourseID     uint
	LessonID     uint
	Completed    bool
	QuizScore    float64 // percentage
	LastPosition int     // seconds into video
	Notes        string  `gorm:"type:text"`
}