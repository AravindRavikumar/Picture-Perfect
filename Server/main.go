package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "./router"
)

func main() {
    router := router.Router()

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./Public/")))
    
    fmt.Println("Server starting at port 3000")
    http.ListenAndServe(":3000", router)
}



