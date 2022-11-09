package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	d,err:= gorm.Open("mysql","ritehsonawane1372:ritesh1372@/mysqlDB?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
		return
	}

}

func GetDB() *gorm.DB{
	return db
}