package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model

	Name string `gorm :"" json:"name"`
	Author string `json :"author"`
	Publication string `json :"publication"`
}


// CreateBook
func CreateBook(book *Book) *Book{

	db.NewRecord(book) 			// Creating new Book table handle by ORM in Gorm
	db.Create(&book)		   // Storing Values
	return book

}

// GetAllBook

func GetAllBook() []Book{
	var books []Book
	db.Find(&books)		// Storing values in books
	return books
}


// GetBookById
func GetBookById( Id int64) (*Book,*gorm.DB){

	var bookById Book
	db := db.Where("Id=?",Id).Find(&bookById)
	return &bookById,db
}

func DeleteBook(Id int64) []Book{

	var bookToDelete Book
	db.Where("Id=?",Id).Delete(&bookToDelete)

	var allBoooks []Book

	db.Find(&allBoooks)
	return allBoooks

}




// Below we initialize the database
func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}
