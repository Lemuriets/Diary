package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Login        string `json:"login" gorm:"not null; unique"`
	PasswordHash string `json:"password" gorm:"not null"`
	SchoolID     uint
	School       School `json:"school" gorm:"foreignkey:SchoolID"`
	ClassID      Class  `json:"class"`
	Permissions  uint8  `json:"permissions"`
}

type Class struct {
	gorm.Model

	ID      uint  `json:"id" gorm:"primarykey:auto_increment"`
	Number  uint8 `json:"number" gorm:"not null"`
	OwnerID uint  `json:"ownerid" gorm:"not null"`
}

type School struct {
	gorm.Model

	ID      uint   `json:"id" gorm:"primarykey:auto_increment"`
	Country string `json:"country" gorm:"not null"`
	Region  string `json:"region" gorm:"not null"`
	ClassID uint
	Classes []Class `json:"classes" gorm:"foreignkey:ClassID"`
}

type Homework struct {
	gorm.Model

	ID       uint   `json:"id" gorm:"primarykey:auto_increment"`
	Payload  string `json:"payload"`
	SchoolID School `json:"school" gorm:"not null"`
	ClassID  Class  `json:"class" gorm:"not null"`
}

type Lesson struct {
	gorm.Model

	Name string `json:"name" gorm:"not null; unique"`
}

type Shedule struct {
	gorm.Model

	ID        uint      `json:"id" gorm:"primarykey:auto_increment"`
	SchoolID  School    `json:"school" gorm:"not null"`
	ClassID   Class     `json:"class" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"not null"`
	DayOfWeek uint8     `json:"dayofweek" gorm:"not null"`
	Lessons   []Lesson  `json:"lessons" gorm:"not null"`
}
