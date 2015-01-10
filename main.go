package main

import (
	"fmt"
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

func startServer() {
	fmt.Println("starting server...")

	http.ListenAndServe(":8080", nil)
}

func main() {
	// db := NewPostgresDB()
	// rce.seedDB(db)
	// rce.RegisterRouteHandlers()
	// startServer()
	fmt.Println("nothing")
}
