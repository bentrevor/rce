package main

import (
	"fmt"
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
	db, err = sql.Open("postgres", "user=rce_admin dbname=rce_dev sslmode=disable")

	if err != nil {
		log.Fatal("failure connecting to database: ", err)
	}
}

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
