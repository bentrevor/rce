package rcfe

// func registerRouteHandlers() {
// 	fmt.Println("registering handlers...")

// 	http.HandleFunc("/", rootHandler)
// 	http.HandleFunc("/hedge_funds", hedgeFundsHandler)
// 	http.HandleFunc("/transaction", transactionHandler)

// 	for _, name := range []string{"hf1", "hf2", "hf3"} {
// 		http.HandleFunc(fmt.Sprintf("/%s", name), hedgeFundHandler(name))
// 	}

// 	http.HandleFunc("/menu/", menuHandler)
// }

// func transactionHandler(w http.ResponseWriter, r *http.Request) {
// 	balances := getBalances()
// 	totalDollars := countTotalDollars(balances)
// 	fmt.Fprintf(w, fmt.Sprintf("dollars before transacting: $%d\n\n", totalDollars))

// 	hf1 := getEntity("hf1", "hedge_funds")
// 	hf2 := getEntity("hf2", "hedge_funds")

// 	traderTransaction := Transaction{
// 		amount:   20,
// 		currency: Dollars,
// 	}
// 	tradeeTransaction := Transaction{
// 		amount:   10,
// 		currency: Pesos,
// 	}
// 	offer := Offer{
// 		traderTransaction: traderTransaction,
// 		tradeeTransaction: tradeeTransaction,
// 	}
// 	trade := Trade{
// 		trader: hf1,
// 		tradee: hf2,
// 		Offer:  offer,
// 		Desc:   "hf1 === 20 Dollars ===> hf2, hf2 === 10 Pesos ===> hf1",
// 	}
// 	trades := []Trade{trade}
// 	transactor := Transactor{}

// 	transactor.ExecuteAll(trades)

// 	totalDollars = countTotalDollars(balances)
// 	fmt.Fprintf(w, fmt.Sprintf("dollars after transacting: $%d", totalDollars))
// }

// func menuHandler(w http.ResponseWriter, r *http.Request) {
// 	menuPage, err := readHTMLFile("menu")

// 	if err != nil {
// 		log.Fatal("there was an error reading the file menu.html")
// 	}

// 	fmt.Fprintf(w, string(menuPage.Body))
// }

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "root")
// }

// func hedgeFundsHandler(w http.ResponseWriter, r *http.Request) {
// 	balances := getBalances()

// 	for name, balance := range balances {
// 		fmt.Fprintf(w, fmt.Sprintf("hedge fund %s has $%d and %d pesos!\n\n", name, balance[Dollars], balance[Pesos]))
// 	}
// }

// func hedgeFundHandler(name string) func(w http.ResponseWriter, r *http.Request) {
// 	if name == "hf1" {
// 		return hf1Handler
// 	} else if name == "hf2" {
// 		return hf2Handler
// 	} else if name == "hf3" {
// 		return hf3Handler
// 	} else {
// 		log.Fatal(fmt.Sprintf("no hedge fund with name %s\n", name))
// 	}

// 	return nil
// }

// func hf1Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, htmlFor("hf1"))
// }

// func hf2Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, htmlFor("hf2"))
// }

// func hf3Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, htmlFor("hf3"))
// }
