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
type Entity interface {
	TableName() string
}

type Bank struct {
	name    string
	dollars int
	pesos   int
}

type HedgeFund struct {
	name    string
	dollars int
	pesos   int
}

type Currency string

const (
	Dollars Currency = "dollars"
	Pesos   Currency = "pesos"
)

func (Bank) TableName() string {
	return "banks"
}

func (HedgeFund) TableName() string {
	return "hedge_funds"
}

// value in Dollars
var exchangeRates = map[Currency]int{
	Dollars: 1,
	Pesos:   2,
}

var db *sql.DB
var err error

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
	Offer  Offer
	Desc   string
}

type Transactor struct{}

func (transactor Transactor) Execute(trade Trade) error {
	fmt.Println(trade.Desc)
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

func getEntity(entityName string, tableName string) Entity {
	query := fmt.Sprintf("select name,dollars,pesos from %s where %s.name = '%s';", tableName, tableName, entityName)
	fmt.Println(fmt.Sprintf("\n\nquery: %s\n\n", query))
	rows, e := db.Query(query)
	defer rows.Close()

	if e != nil {
		log.Fatal("failure querying the database: ", e)
	}

	// balances := getBalances()

	return &HedgeFund{}
}

func readHTMLFile(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func htmlFor(filename string) string {
	page, err := readHTMLFile(filename)

	if err != nil {
		log.Fatal(fmt.Sprintf("there was an error reading the file %s.html\n", filename))
	}

	return string(page.Body)
}

// TODO make a data structure for this
func getBalances() map[string]map[Currency]int {
	rows, e := db.Query("select name,dollars,pesos from hedge_funds;")
	defer rows.Close()

	if e != nil {
		log.Fatal("failure querying the database: ", e)
	}

	balances := map[string]map[Currency]int{}

	for rows.Next() {
		var name string
		var dollars int
		var pesos int
		rows.Scan(&name, &dollars, &pesos)

		balances[name][Dollars] = dollars
		balances[name][Pesos] = pesos
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
  dollars integer,
  pesos   integer
);

insert into hedge_funds (name, dollars, pesos) values ('hf1', 1234, 123),
                                                      ('hf2', 5678, 567),
                                                      ('hf3', 9090, 909)
;

drop table if exists banks;
create table banks (
  id      serial,
  name    varchar(50),
  dollars integer,
  pesos   integer
);

insert into banks (name, dollars, pesos) values ('b1', 11111, 1111),
                                                ('b2', 22222, 2222),
                                                ('b3', 33333, 3333)
;
`)

	if err != nil {
		log.Fatal("error seeding the database: ", err)
	}
}

func main() {
	connectToDb()
	seedDB()
	// rcfe.RegisterRouteHandlers()
	startServer()
}
