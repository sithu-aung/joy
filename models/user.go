package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Email string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	FirstName string
	LastName string
	Role string `gorm:"default:student"`
	ProfilePic string
	Bio string
	Courses []Course `gorm:"many2many:user_courses"`
	Progress []UserProgress
}