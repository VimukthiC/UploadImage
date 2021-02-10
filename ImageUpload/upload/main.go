package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func run(w http.ResponseWriter , r *http.Request)  {
	template1,err:=template.ParseFiles("index.html")
	if err !=nil{
		panic(err.Error())
	}
	err=template1.Execute(w,nil)
	if err!=nil{
		panic(err.Error())
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",run)
	http.ListenAndServe(":3000",router)
}
