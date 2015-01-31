package trie

import "sort"

// Contains returns true if word exists in trie
func (tr Trie) Contains(word string) bool {
	r := tr.root
	for _, b := range []byte(word) {
		if r.children[b-offset] == nil {
			return false
		}
		r = r.children[b-offset]
	}
	return r.final
}

// List returns []string containing all words held in dictionary
//
func (tr Trie) List() []string {
	return tr.ListFrom("")
}

// ListFrom returns a list of words in dictionary starting with the prefix
// If prefix doesn't exist, returns empty list
//
func (tr Trie) ListFrom(prefix string) []string {
	st := tr.root
	for _, b := range []byte(prefix) {
		b -= offset
		if st.children[b] == nil {
			return []string{}
		}
		st = st.children[b]
	}

	stream := make(chan string, 1000)

	go func() {
		traverse(st, []byte(prefix), stream)
		close(stream)
	}()

	l := []string{}
	for word := range stream {
		l = append(l, word)
	}
	sort.Strings(l)
	return l
}

func traverse(n *state, prefix []byte, stream chan string) {
	for i, n := range n.children {
		if n != nil {
			traverse(n, append(prefix, byte(i+offset)), stream)
		}
	}
	if n.final {
		stream <- string(prefix)
	}
}
