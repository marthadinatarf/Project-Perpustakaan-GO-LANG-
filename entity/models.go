package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Gender   string
	Nohp     string
}

type Book struct {
	gorm.Model
	Judul    string
	Author   string
	Penerbit string
}
