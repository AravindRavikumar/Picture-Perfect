package middleware

import (
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

func CatalogueHandler(w http.ResponseWriter, r *http.Request) {

}

/*
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    //push(w, "./Public/css/style.css")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    homeTemplate, err := template.ParseFiles("./Components/Home/home.html")
    if err != nil {
		fmt.Println(err)
	}
	homeTemplate.Execute(w, nil)
}
*/