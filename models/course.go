package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Category    string
	Level       string // beginner, intermediate, advanced
	ImageURL    string
	AuthorID    uint
	Lessons     []Lesson
	Students    []User `gorm:"many2many:user_courses;"`
}

// Lesson represents a lesson within a course
type Lesson struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Content     string `gorm:"type:text"`
	VideoURL    string
	Order       int
	CourseID    uint
	Duration    int // in minutes
	Quiz        *Quiz
}