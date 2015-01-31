package dawg

import "sort"

// List emits all words contained in the dictionary, in order
//
func (dg Dawg) List() []string {
	return dg.ListFrom("")
}

// ListFrom returns all words in the dictionary starting with the given prefix.
// If a prefix does not exist, returns an empty list.
func (dg Dawg) ListFrom(prefix string) []string {
	i, st := dg.root.getPrefix(prefix)
	if prefix != prefix[:i] {
		return []string{}
	}

	list := []string{}
	stream := make(chan string, 1000)
	go func() {
		traverse(st, []byte(prefix), stream)
		close(stream)
	}()

	for word := range stream {
		list = append(list, word)
	}
	sort.Strings(list)
	return list
}

func traverse(st *State, prefix []byte, stream chan string) {
	for i, st := range st.children {
		if st != nil {
			traverse(st, append(prefix, byte(i+offset)), stream)
		}
	}
	if st.final {
		stream <- string(prefix)
	}
}
