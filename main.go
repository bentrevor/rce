package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"database/sql"

	_ "github.com/lib/pg"
)

type Page struct {
	Title string
	Body  []byte
}

func readHTMLFile(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	menuPage, err := readHTMLFile("menu")

	if err != nil {
		panic("there was an error reading the file menu.html")
	}

	fmt.Fprintf(w, string(menuPage.Body))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root")
}

func main() {
	db, _ := sql.Open("postgres", "user=rcse_admin dbname=rcse_dev")
	fmt.Println("starting server...\n")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/menu/", menuHandler)

	http.ListenAndServe(":8080", nil)
}
