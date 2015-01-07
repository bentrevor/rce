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

func getEntity(name string, tableName string) Entity {
	rows, e := db.Query(fmt.Sprintf("select name,dollars,pesos from %s where %s.name = %s;", tableName, tableName, name))
	defer rows.Close()

	if e != nil {
		log.Fatal("failure querying the database: ", e)
	}

	balances := map[string]int{}

	for rows.Next() {
		var dollars int
		var pesos int
		var name string
		rows.Scan(&name, &dollars, &pesos)
		balances[name] = dollars
	}

	return &HedgeFund{}
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	balances := getBalances()
	totalDollars := countTotalDollars(balances)
	fmt.Fprintf(w, fmt.Sprintf("dollars before transacting: $%d\n\n", totalDollars))

	hf1 := getEntity("hf1", "hedge_funds")
	hf2 := getEntity("hf2", "hedge_funds")

	traderTransaction := Transaction{
		amount:   20,
		currency: Dollars,
	}
	tradeeTransaction := Transaction{
		amount:   10,
		currency: Pesos,
	}
	offer := Offer{
		traderTransaction: traderTransaction,
		tradeeTransaction: tradeeTransaction,
	}
	trade := Trade{
		trader: hf1,
		tradee: hf2,
		offer:  offer,
	}
	trades := []Trade{trade}
	transactor := Transactor{}

	transactor.ExecuteAll(trades)

	totalDollars = countTotalDollars(balances)
	fmt.Fprintf(w, fmt.Sprintf("dollars after transacting: $%d", totalDollars))
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

func hedgeFundsHandler(w http.ResponseWriter, r *http.Request) {
	balances := getBalances()

	for name, dollars := range balances {
		fmt.Fprintf(w, fmt.Sprintf("hedge fund %s has $%d!\n\n", name, dollars))
	}
}

func htmlFor(filename string) string {
	page, err := readHTMLFile(filename)

	if err != nil {
		log.Fatal(fmt.Sprintf("there was an error reading the file %s.html\n", filename))
	}

	return string(page.Body)
}

func hf1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlFor("hf1"))
}

func hf2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlFor("hf2"))
}

func hf3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlFor("hf3"))
}

func hedgeFundHandler(name string) func(w http.ResponseWriter, r *http.Request) {
	if name == "hf1" {
		return hf1Handler
	} else if name == "hf2" {
		return hf2Handler
	} else if name == "hf3" {
		return hf3Handler
	} else {
		log.Fatal(fmt.Sprintf("no hedge fund with name %s\n", name))
	}

	return nil
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
	// TODO duplication
	http.HandleFunc("/hf1", hf1Handler)
	http.HandleFunc("/hf2", hedgeFundHandler("hf2"))
	http.HandleFunc("/hf3", hedgeFundHandler("hf3"))
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
	registerRouteHandlers()
	startServer()
}
