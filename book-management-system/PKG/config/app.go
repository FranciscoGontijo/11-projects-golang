package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect(){
	d, err := gorm.Open("mysql", "francisco:Franciscotest/bookstore?charset=utf8&parseTime=True&loc=local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}