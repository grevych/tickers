package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/grevych/tickers"
)

type fetchedTicker struct {
	Ticker   string `json:"ticker"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
}

func search(trie tickers.Trie) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		symbol := ""
		if queryParameter, ok := r.URL.Query()["symbol"]; ok {
			symbol = queryParameter[0]
		}

		results := trie.Search(symbol)
		response, err := json.Marshal(results)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	}
}

func fetch() []*fetchedTicker {
	URL := "https://dumbstockapi.com/stock"
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("error while fetching stock")
	}
	defer resp.Body.Close()

	var fetchedTickers []*fetchedTicker
	if err := json.NewDecoder(resp.Body).Decode(&fetchedTickers); err != nil {
		log.Fatal("error while parsing tickers")
	}

	return fetchedTickers
}

func main() {
	trie := tickers.NewTrie()

	for _, ticker := range fetch() {
		trie.Add(
			*tickers.NewTicker(ticker.Ticker, ticker.Name, ticker.Exchange),
		)
	}

	router := mux.NewRouter()
	handler := search(trie)
	router.HandleFunc("/search", handler).Methods("GET")

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(address, router))
}
