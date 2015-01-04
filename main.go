package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/lib/pq"
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
		log.Fatal("there was an error reading the file menu.html")
	}

	fmt.Fprintf(w, string(menuPage.Body))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root")
}

func main() {
	fmt.Println("connecting to db...")
	// TODO don't disable ssl...
	db, err := sql.Open("postgres", "user=rcfe_admin dbname=rcfe_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}

	fmt.Println("seeding db...")
	seedDB(db)

	rows, e := db.Query("select name,dollars from hedge_funds;")
	defer rows.Close()

	if e != nil {
		log.Fatal("failure querying the database: ", e)
	}

	for rows.Next() {
		var dollars int
		var name string
		rows.Scan(&name, &dollars)
		fmt.Printf("found the hedge fund %s with $%d\n\n", name, dollars)
	}

	fmt.Println("registering handlers...")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/menu/", menuHandler)

	fmt.Println("starting server...")

	http.ListenAndServe(":8080", nil)
}

func seedDB(db *sql.DB) {
	_, err := db.Query(`
drop table if exists hedge_funds;
create table hedge_funds (
  id      serial,
  name    varchar(50),
  dollars integer
);

insert into hedge_funds (name, dollars) values ('asdf', 123);

drop table if exists banks;
create table banks (
  id      serial,
  name    varchar(50),
  dollars integer
);
`)

	if err != nil {
		log.Fatal("error seeding the database: ", err)
	}
}
