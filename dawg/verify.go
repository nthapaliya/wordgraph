package dawg

import "fmt"

// Verify checks that there are no redundant nodes, and that the pruning
// by the algorithm is successful.
//
func (dg *Dawg) Verify() error {
	register := dg.register
	seen := make(map[string]bool)

	stream := make(chan *State, 1000)
	go func() {
		st := dg.root
		_start(st, stream)
		close(stream)
	}()

	redundant := 0
	for st := range stream {
		hash := st.hash
		if _, ok := register[hash]; !ok && st != dg.root {
			// return false, errors.New("Dawg not minimized")
			if !seen[hash] {
				redundant++
				seen[hash] = true
			}
		}
	}
	if redundant == 0 {
		return nil
	}
	return fmt.Errorf("Dawg not minimal. %d unregistered/redundant nodes", redundant)
}

func _start(st *State, stream chan *State) {
	stream <- st
	for _, c := range st.children {
		if c != nil {
			_start(c, stream)
		}
	}
}
