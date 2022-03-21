package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Login        string `json:"login" gorm:"not null; unique; size:100"`
	PasswordHash string `json:"password" gorm:"not null"`
	Name         string `json:"name" gorm:"not null"`
	Lastname     string `json:"lastname" gorm:"not null"`
	Patronymic   string `json:"Patronymic" gorm:"not null"`
	ClassID      uint   `json:"classId"`
	Class        Class  `json:"class" gorm:"foreignkey:ClassID"`
	Permissions  uint8  `json:"permissions"`
}

type Class struct {
	gorm.Model

	ID       uint   `json:"id" gorm:"primarykey:autoIncrement"`
	Number   uint8  `json:"number" gorm:"not null"`
	OwnerID  uint   `json:"ownerId" gorm:"not null"`
	SchoolID uint   `json:"SchoolID"`
	School   School `json:"schoolId" gorm:"foreignkey:SchoolID"`
}

type School struct {
	gorm.Model

	ID      uint   `json:"id" gorm:"primarykey:autoIncrement"`
	Country string `json:"country" gorm:"not null"`
	Region  string `json:"region" gorm:"not null"`
}

type Homework struct {
	gorm.Model

	ID      uint   `json:"id" gorm:"primarykey:autoIncrement"`
	Payload string `json:"payload"`
	ClassID uint   `json:"classId"`
	Class   Class  `json:"class" gorm:"not null; foreignkey:ClassID"`
}

type Lesson struct {
	gorm.Model

	ID        uint    `json:"id" gorm:"primarykey:autoIncrement"`
	Name      string  `json:"name" gorm:"not null; unique"`
	SheduleID uint    `json:"sheduleId"`
	Shedule   Shedule `json:"shedule" gorm:"foreignkey:SheduleID"`
}

type Shedule struct {
	gorm.Model

	ID        uint      `json:"id" gorm:"primarykey:autoIncrement"`
	Date      time.Time `json:"date" gorm:"not null"`
	DayOfWeek uint8     `json:"dayOfWeek" gorm:"not null"`
	ClassID   uint      `json:"classId"`
	Class     Class     `json:"class" gorm:"not null; foreignkey:ClassID"`
}

type Mark struct {
	gorm.Model

	Number   uint8  `json:"number"`
	UserID   uint   `json:"userId"`
	User     User   `json:"user" gorm:"foreignkey:UserID"`
	LessonID uint   `json:"lessonId"`
	Lesson   Lesson `json:"lesson" gorm:"foreignkey:LessonID"`
}
