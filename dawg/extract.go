package dawg

import "sort"

// List returns words contained by the dictionary in sorted order
func (dg Dawg) List() []string {
	list := []string{}
	stream := make(chan string, 1000)
	go func() {
		traverse(dg.root, []byte{}, stream)
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
