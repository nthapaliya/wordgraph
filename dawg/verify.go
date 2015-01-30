package dawg

import (
	"fmt"
	"sort"
	"strings"
)

// Verify checks that there are no redundant nodes, and that the pruning
// by the algorithm is successful.
func (dg *Dawg) Verify() (bool, error) {
	register := dg.register

	stream := make(chan *State, 1000)
	go func() {
		st := dg.root
		_start(st, stream)
		close(stream)
	}()

	for st := range stream {
		hash := st.hash
		if _, ok := register[hash]; !ok {
			if st != dg.root {
				return false, fmt.Errorf("error")
			}
		}
	}
	return true, nil
}

func _start(st *State, stream chan *State) {
	stream <- st
	for _, c := range st.children {
		if c != nil {
			_start(c, stream)
		}
	}
}

func sortAndToLower(l []string) []string {
	if !sort.StringsAreSorted(l) {
		sort.Strings(l)
	}
	for i, word := range l {
		l[i] = strings.ToLower(word)
	}
	return l
}
