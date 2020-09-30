package tickers

// Ticker gathers the description of a ticker
type Ticker struct {
	Symbol     string
	Exchange   string
	Instrument string
}

// NewTicker creates a pointer to a ticker
func NewTicker(symbol, exchange, instrument string) *Ticker {
	return &Ticker{symbol, exchange, instrument}
}
