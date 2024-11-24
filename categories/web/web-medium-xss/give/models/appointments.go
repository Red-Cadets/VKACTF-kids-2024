package models

import (
	"gorm.io/gorm"
)

type Appointments struct {
	gorm.Model `gorm:"-"`
	Login      string `gorm:"column:Login;type:varchar(124);not null"`
	Date       string `gorm:"column:Date;type:varchar(20);not null"`
	Time       string `gorm:"column:Time;type:varchar(20);not null"`
	Doctor     string `gorm:"column:Doctor;type:varchar(20);not null"`
	Complain   string `gorm:"column:Complain;type:varchar(1000);not null"`
	EncodedURL string `gorm:"column:EncodedURL;type:varchar(255);not null"`
}
