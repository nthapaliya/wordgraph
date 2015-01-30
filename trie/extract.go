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
func (tr Trie) List() []string {
	stream := make(chan string, 1000)

	go func() {
		traverse(tr.root, []byte{}, stream)
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
