package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// gather URL vars from mux
	v := mux.Vars(r)

	title, ok := v["title"]
	if !ok {
		title = "This"
	}

	// build Page struct for template
	p := &Page{Title: title, Body: "This & that"}

	// load template
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	// build template
	err = t.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("Setting up routes ... \n")
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/{title}", IndexHandler)
	http.Handle("/", r)

	fmt.Printf("Starting http Server ...\n")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
