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
	Age          uint8  `json:"age"`
	ClassID      uint64 `json:"classId"`
	Class        Class  `json:"class" gorm:"foreignkey:ClassID"`
	Permissions  uint8  `json:"permissions"`
}

type Class struct {
	gorm.Model

	Number   uint8  `json:"number" gorm:"not null"`
	OwnerID  uint64 `json:"ownerId" gorm:"not null"`
	SchoolID uint64 `json:"SchoolID"`
	School   School `json:"school" gorm:"foreignkey:SchoolID"`
}

type School struct {
	gorm.Model

	Country string `json:"country" gorm:"not null"`
	Region  string `json:"region" gorm:"not null"`
}

type Homework struct {
	gorm.Model

	Payload string `json:"payload"`
	ClassID uint64 `json:"classId"`
	Class   Class  `json:"class" gorm:"not null; foreignkey:ClassID"`
}

type Lesson struct {
	gorm.Model

	Name      string  `json:"name" gorm:"not null; unique"`
	SheduleID uint64  `json:"sheduleId"`
	Shedule   Shedule `json:"shedule" gorm:"foreignkey:SheduleID"`
}

type Shedule struct {
	gorm.Model

	Date      time.Time `json:"date" gorm:"not null"`
	DayOfWeek uint8     `json:"dayOfWeek" gorm:"not null"`
	ClassID   uint64    `json:"classId"`
	Class     Class     `json:"class" gorm:"not null; foreignkey:ClassID"`
}

type Mark struct {
	gorm.Model

	Number   uint8  `json:"number"`
	UserID   uint64 `json:"userId"`
	User     User   `json:"user" gorm:"foreignkey:UserID"`
	LessonID uint64 `json:"lessonId"`
	Lesson   Lesson `json:"lesson" gorm:"foreignkey:LessonID"`
}

type RefreshSession struct {
	gorm.Model

	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
}
