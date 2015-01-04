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

var db *sql.DB
var err error

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

func hedgeFundsHandler(w http.ResponseWriter, r *http.Request) {
	rows, e := db.Query("select name,dollars from hedge_funds;")
	defer rows.Close()

	if e != nil {
		log.Fatal("failure querying the database: ", e)
	}

	for rows.Next() {
		var dollars int
		var name string
		rows.Scan(&name, &dollars)
		fmt.Fprintf(w, fmt.Sprintf("hedge fund %s has $%d!\n\n", name, dollars))
	}
}

func connectToDb() {
	fmt.Println("connecting to db...")
	// TODO don't disable ssl...
	db, err = sql.Open("postgres", "user=rcfe_admin dbname=rcfe_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}
}

func routeHandlers() {
	fmt.Println("registering handlers...")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hedge_funds", hedgeFundsHandler)
	http.HandleFunc("/menu/", menuHandler)
}

func startServer() {
	fmt.Println("starting server...")

	http.ListenAndServe(":8080", nil)
}

func seedDB() {
	fmt.Println("seeding db...")

	_, err := db.Query(`
drop table if exists hedge_funds;
create table hedge_funds (
  id      serial,
  name    varchar(50),
  dollars integer
);

insert into hedge_funds (name, dollars) values ('hf1', 1234),
                                               ('hf2', 5678),
                                               ('hf3', 9090)
;

drop table if exists banks;
create table banks (
  id      serial,
  name    varchar(50),
  dollars integer
);

insert into banks (name, dollars) values ('b1', 11111),
                                         ('b2', 22222),
                                         ('b3', 33333)
;
`)

	if err != nil {
		log.Fatal("error seeding the database: ", err)
	}
}

func main() {
	connectToDb()
	seedDB()
	routeHandlers()
	startServer()
}
