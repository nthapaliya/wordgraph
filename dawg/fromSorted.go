package dawg

import (
	"errors"
	"sort"
)

const offset = 97

// NewFromList creates a new Dawg from a slice of sorted strings
func NewFromList(wordlist []string) (*Dawg, error) {
	if !sort.StringsAreSorted(wordlist) {
		return nil, errors.New("Input is not sorted")
	}

	dg := &Dawg{
		root:     &State{false, &Child{}, ""},
		register: make(map[string]*State),
	}

	for _, word := range wordlist {
		index, prefix := dg.root.getPrefix(word)

		if hasChildren(prefix) {
			dg.replaceOrRegister(prefix)
		}
		prefix.addSuffix(word[index:])
	}
	dg.replaceOrRegister(dg.root)

	// Final Cleanup
	return dg, nil
}

func (dg *Dawg) replaceOrRegister(st *State) {
	lastchild := lastState(st)
	if hasChildren(lastchild) {
		dg.replaceOrRegister(lastchild)
	}

	hash := lastchild.getHash()
	if v, ok := dg.register[hash]; ok {
		st.children[lastIndex(st)] = v
	} else {
		dg.register[hash] = lastchild
	}
}
