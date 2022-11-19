package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

// *httpRequest bcoz we will receive something
func GetBook(w http.ResponseWriter, r *http.Request) {

	newBook := models.GetAllBook()

	res, _ := json.Marshal(newBook) // Making the json

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // This will Send the reponse back to the postman

}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) // This will get the res from user
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) // Converting the string into int

	if err != nil {
		log.Fatal("An Error Occured")
	}

	booksDetails, _ := models.GetBookById(ID)

	// res to user
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)
	b:= models.CreateBook(CreateBook)

	res,_:=json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func DeleteBook (w http.ResponseWriter,r * http.Request){

	vars:= mux.Vars(r)

	bookId:=vars["bookId"]

	ID,err:= strconv.ParseInt(bookId,0,0)

	if err!=nil{
		fmt.Println(err)
		log.Fatal("An Error Ocuured")
	}

	book:= models.DeleteBook(ID)

	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


// 1:02:46 Update Route Is remaining 