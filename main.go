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

// TODO better name for (banks and hedge funds)
type Entity struct {
	name    string
	dollars int
}

type Currency struct {
	name           string
	valueInDollars int
}

type Transaction struct {
	amount   int
	currency Currency
}

type Offer struct {
	traderTransaction Transaction
	tradeeTransaction Transaction
}

// TODO trader/tradee look too similar...
type Trade struct {
	trader Entity
	tradee Entity
	offer  Offer
}

type Transactor struct{}

func (transactor Transactor) Execute(trade Trade) error {
	return nil
}

func (transactor Transactor) ExecuteAll(trades []Trade) error {
	for _, trade := range trades {
		err := transactor.Execute(trade)

		if err != nil {
			log.Fatal("error in ExecuteAll: ", err)
		}
	}

	return nil
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	balances := getBalances()
	totalDollars := countTotalDollars(balances)
	fmt.Fprintf(w, fmt.Sprintf("dollars before transacting: $%d\n\n", totalDollars))

	trades := []Trade{}
	transactor := Transactor{}

	transactor.ExecuteAll(trades)

	totalDollars = countTotalDollars(balances)
	fmt.Fprintf(w, fmt.Sprintf("dollars after transacting: $%d", totalDollars))
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
	balances := getBalances()

	for name, dollars := range balances {
		fmt.Fprintf(w, fmt.Sprintf("hedge fund %s has $%d!\n\n", name, dollars))
	}
}

func getBalances() map[string]int {
	rows, e := db.Query("select name,dollars from hedge_funds;")
	defer rows.Close()

	if e != nil {
		log.Fatal("failure querying the database: ", e)
	}

	balances := map[string]int{}

	for rows.Next() {
		var dollars int
		var name string
		rows.Scan(&name, &dollars)
		balances[name] = dollars
	}

	return balances
}

func countTotalDollars(balances map[string]int) int {
	total := 0

	for _, b := range balances {
		total += b
	}

	return total
}

func connectToDb() {
	fmt.Println("connecting to db...")
	// TODO don't disable ssl...
	db, err = sql.Open("postgres", "user=rcfe_admin dbname=rcfe_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}
}

func registerRouteHandlers() {
	fmt.Println("registering handlers...")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hedge_funds", hedgeFundsHandler)
	http.HandleFunc("/transaction", transactionHandler)
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
	registerRouteHandlers()
	startServer()
}
