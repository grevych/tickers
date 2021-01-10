package tickers

// Trie is the main data structure that manages tickers
type Trie interface {
	Add(Ticker)
	Size() int
	Search(string) []Ticker
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

func (t *trie) Add(ticker Ticker) {
	if ticker.Symbol == "" {
		return
	}

	cur := t.root
	for _, key := range ticker.Symbol {
		if !cur.hasChild(key) {
			cur.createChild(key)
		}

		cur, _ = cur.getChild(key)
	}

	cur.addTicker(&ticker)
	t.size++
}

func (t *trie) Search(input string) []Ticker {
	cur := t.root

	for _, key := range input {
		if !cur.hasChild(key) {
			return []Ticker{}
		}

		cur, _ = cur.getChild(key)
	}

	response := cur.getTickers()
	queue := append([]*node{}, cur.getChildren()...)
	for len(queue) > 0 {
		// Pop Front
		child := queue[0]
		queue = queue[1:]

		response = append(response, child.getTickers()...)
		// Push Front
		queue = append(child.getChildren(), queue...)
	}

	return response
}

func (t *trie) Size() int {
	return t.size
}
