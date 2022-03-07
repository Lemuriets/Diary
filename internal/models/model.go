package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Login        string `json:"login"`
	PasswordHash string `json:"password"`
	School       uint   `json:"school"`
	Class        string
}

type School struct {
	gorm.Model

	ID      uint   `json:"id",gorm:"PrimaryKey"`
	Country string `json:"country"`
	Region  string `json:"region"`
}

type Class struct {
	gorm.Model

	ID      uint `json:"id",gorm:"PrimaryKey"`
	OwnerID uint
}

type Homework struct {
	gorm.Model

	ID uint `json:"id",gorm:"PrimaryKey"`
}

type Shedule struct {
	gorm.Model

	School string
	Class  string
}
