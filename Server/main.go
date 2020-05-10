package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "github.com/unrolled/render"
)
 
func main() {
    renderObj := render.New()
    router := mux.NewRouter()
    
    router.HandleFunc("/catalogue", CatalogueHandler)
    router.HandleFunc("/home", HomeHandler)
    http.ListenAndServe(":3000", router)
}
 
func CatalogueHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    page := vars["page"]

    fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    
}