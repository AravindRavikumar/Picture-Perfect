package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/catalogue", CatalogueHandler)
    router.HandleFunc("/home", HomeHandler)
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./Public/")))
    
    http.ListenAndServe(":3000", router)
}

/*
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			fmt.Printf("Failed to push: %v", err)
            return
		}
	}
}
*/

func CatalogueHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    page := vars["page"]

    fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    //push(w, "./Public/css/style.css")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    homeTemplate, err := template.ParseFiles("./Components/Home/home.html")
    if err != nil {
		fmt.Println(err)
	}
	homeTemplate.Execute(w, nil)
}

