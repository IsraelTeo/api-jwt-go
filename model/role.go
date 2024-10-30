package model

import "gorm.io/gorm"

type Users []User

type Role struct {
	gorm.Model

	ID     uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Name   string `json:"name" gorm:"unique;not null"`
	Status bool   `json:"status" gorm:"defaut:true"`
	Users  Users  `json:"users"`
}
