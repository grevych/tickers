package tickers

// Trie is the main data structure that manages tickers
type Trie interface {
	next(*node, rune) *node
	add(Ticker)
	GetSize() int
	LoadTickers([]Ticker)
	GetTickers(string) []*Ticker
}

type trie struct {
	root *node
	size int
}

var _ Trie = &trie{}

// NewTrie creates a trie in form of an interface
func NewTrie() Trie {
	return &trie{
		root: newNode(),
	}
}

func (t *trie) add(ticker Ticker) {
	if ticker.Symbol == "" {
		return
	}

	ptr := t.root
	aux := ptr
	for _, r := range ticker.Symbol {
		if !ptr.hasChild(r) {
			child := newNode()
			ptr.addChild(r, child)
		}

		aux = ptr
		ptr, _ = ptr.getChild(r)
		// aux.print()
	}

	aux.addTicker(&ticker)
	// aux.print()
	// fmt.Println("")
	t.size++
}

func (t *trie) next(ptr *node, r rune) *node {
	if !ptr.hasChild(r) {
		return nil
	}

	child, _ := ptr.getChild(r)
	// child.print()

	return child
}

func (t *trie) GetTickers(input string) []*Ticker {
	ptr := t.root
	aux := ptr

	for _, r := range input {
		aux = ptr
		ptr = t.next(ptr, r)

		if ptr == nil {
			return []*Ticker{}
		}
	}

	return append(aux.tickers, ptr.getChildrenTickers()...)
}

func (t *trie) LoadTickers(tickers []Ticker) {
	for _, tk := range tickers {
		t.add(tk)
	}
}

func (t *trie) GetSize() int {
	return t.size
}
