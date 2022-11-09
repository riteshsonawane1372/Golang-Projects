package main

import(
	"fmt"
	"log"
	"net/http"
)


func main(){
	fileserver:=http.FileServer(http.Dir("./static")) 	// Telling Static has html

	
	http.Handle("/",fileserver)	// Telling that / route will be handle by fileserver
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf(" \n Starting At Port 8080: ")

	if err:= http.ListenAndServe(":8080",nil); err!=nil{

		log.Fatal(err)
	}

}

// w is the response and r is the request
func helloHandler(w http.ResponseWriter,r *http.Request){
	
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method !="GET"{
		http.Error(w,"method is not supported",http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"hello")
}

func formHandler(w http.ResponseWriter,r *http.Request){
	if err:= r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParseForm err",err)
		return
	}
	fmt.Fprintf(w,"POST request successful")
	name:=r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w,"Name %s\n",name)
	fmt.Fprintf(w,"address %s\n",address)

}

