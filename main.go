package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func yourBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t := template.Must(template.ParseFiles("index.html"))

	books := map[string][]Book{
		"Books": {
			{"The Alchemist", "Paulo Coelho", 1988},
			{"Atomic Habits", "James Clear", 2018},
			{"Thinking Fast and Slow", "Daniel Kahneman", 2011},
			{"Deep Work", "Cal Newport", 2016},
			{"The Book of Five Rings", "Miyamoto Musashi", 1645},
		},
	}

	t.Execute(w, books)
	fmt.Println(r.Method)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	r.ParseForm()
	//t := template.Must(template.ParseFiles("index.html"))

	title := r.PostFormValue("title")
	author := r.PostFormValue("author")

	year := r.PostFormValue("year")

	htmlStr := fmt.Sprintf("<li class='list-group-item text-bg-light p-3 text-dark'>%s - %s - %s - <button type='button' class='btn-close' aria-label='Close'></button></li>", title, author, year)

	t, _ := template.New("t").Parse(htmlStr)

	// t.ExecuteTemplate(w, "film-list-element", Book{Title: title, Author: author, Year: year})

	t.Execute(w, nil)
	fmt.Println(r.Method)
}

func main() {
	http.HandleFunc("/", yourBook)
	http.HandleFunc("/add-book/", addBook)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
