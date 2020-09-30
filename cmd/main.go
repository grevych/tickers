package main

import (
	"fmt"

	"github.com/grevych/tickers"
)

func printTickers(key string, tickers []*tickers.Ticker) {
	fmt.Printf("%s -> ", key)
	for _, ticker := range tickers {
		fmt.Printf("%v", ticker)
	}
	fmt.Println("")
}

func main() {
	trie := tickers.NewTrie()

	trie.LoadTickers([]tickers.Ticker{
		*tickers.NewTicker("A", "NASDAQ", ""),
		*tickers.NewTicker("AA", "NASDAQ", ""),
		*tickers.NewTicker("AAPL", "NASDAQ", ""),
		*tickers.NewTicker("AA", "BMV", ""),
		*tickers.NewTicker("ABA", "BMV", ""),
		*tickers.NewTicker("NET", "NASDAQ", ""),
		*tickers.NewTicker("GOOG", "NASDAQ", ""),
		*tickers.NewTicker("GOOGL", "NASDAQ", ""),
		*tickers.NewTicker("MSFT", "NASDAQ", ""),
		*tickers.NewTicker("FB", "NASDAQ", ""),
		*tickers.NewTicker("NVDA", "NASDAQ", ""),
		*tickers.NewTicker("AMZN", "NASDAQ", ""),
	})

	printTickers("A", trie.GetTickers("A"))
	printTickers("AA", trie.GetTickers("AA"))
	printTickers("GO", trie.GetTickers("GO"))
}
