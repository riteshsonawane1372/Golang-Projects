package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// This File Will Return a DB

var (
	db *gorm.DB
)

func Connect(){

	d,err := gorm.Open("mysql","ritesh:Ritesh1372/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err !=nil{
		fmt.Println(err)
		panic(err)
	}

	db=d

}

func GetDB() *gorm.DB{
	return db
}