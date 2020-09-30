package tickers

import "fmt"

type node struct {
	children map[rune]*node
	tickers  []*Ticker
}

func newNode() *node {
	return &node{
		children: map[rune]*node{},
		tickers:  []*Ticker{},
	}
}

func (n *node) addChild(key rune, child *node) error {
	if _, exists := n.children[key]; exists {
		return fmt.Errorf("Node error: child %c exists", key)
	}

	n.children[key] = child

	return nil
}

func (n *node) getChild(key rune) (*node, error) {
	child, exists := n.children[key]

	if !exists {
		return nil, fmt.Errorf("Node error: child doesn't %c exists", key)
	}

	return child, nil
}

func (n *node) hasChild(key rune) bool {
	_, exists := n.children[key]

	return exists
}

func (n *node) addTicker(ticker *Ticker) {
	n.tickers = append(n.tickers, ticker)
}

func (n *node) getTickers() []*Ticker {
	return n.tickers
}

func (n *node) getChildrenTickers() []*Ticker {
	childrenTickers := append([]*Ticker{}, n.tickers...)

	for _, child := range n.children {
		childrenTickers = append(childrenTickers, child.getChildrenTickers()...)
	}

	return childrenTickers
}

func (n *node) print() {
	keys := []string{}
	for key, _ := range n.children {
		keys = append(keys, string(key))
	}

	tickers := []Ticker{}
	for _, t := range n.tickers {
		tickers = append(tickers, *t)
	}

	fmt.Printf("Keys: %v, Tickers: %v \n", keys, tickers)
}
