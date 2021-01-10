package tickers

// Ticker gathers the description of a ticker
type Ticker struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
}

// NewTicker creates a pointer to a ticker
func NewTicker(symbol, name, exchange string) *Ticker {
	return &Ticker{symbol, name, exchange}
}
