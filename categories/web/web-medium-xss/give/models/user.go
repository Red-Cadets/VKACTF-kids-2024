package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model     `gorm:"-"`
	Name           string `gorm:"column:Name;type:varchar(255);not null"`
	Surname        string `gorm:"column:Surname;type:varchar(255);not null"`
	DateOfBirthday string `gorm:"column:DateOfBirthday;type:varchar(255);not null"`
	Email          string `gorm:"column:Email;type:varchar(255);not null"`
	PhoneNumber    string `gorm:"column:PhoneNumber;type:varchar(255);not null"`
	Passport       string `gorm:"column:Passport;type:varchar(255);not null"`
	Login          string `gorm:"column:Login;type:varchar(124);not null"`
	Password       string `gorm:"-"`
	PassHash       string `gorm:"column:PassHash;type:varchar(255);not null"`
}
