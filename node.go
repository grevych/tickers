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

func (n *node) createChild(key rune) error {
	if _, exists := n.children[key]; exists {
		return fmt.Errorf("Node error: child %c exists", key)
	}

	n.children[key] = newNode()

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

func (n *node) getTickers() []Ticker {
	tickers := []Ticker{}
	for _, ticker := range n.tickers {
		tickers = append(tickers, *ticker)
	}

	return tickers
}

func (n *node) getChildren() []*node {
	children := []*node{}

	for _, value := range n.children {
		children = append(children, value)
	}

	return children
}
