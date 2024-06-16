package model

import (
	"gorm.io/gorm"
)

type User1 struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

// 定义一个结构体，用于存储token
type Token struct {
	Token string `json:"token"`
}
