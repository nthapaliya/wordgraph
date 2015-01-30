package trie

type Trie struct {
	root *state
}

type state struct {
	final    bool
	children *child
	hash     string
}

type child [26]*state
