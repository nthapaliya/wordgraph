package dawg

import (
	"container/list"
	"errors"
	"sort"
)

// NewFromList creates a new Dawg from a slice of sorted strings
//
func NewFromList(wordlist []string) (*Dawg, error) {
	if !sort.StringsAreSorted(wordlist) {
		return nil, errors.New("Unsorted list")
	}

	dg := &Dawg{
		root:     newState(0),
		register: make(map[string]*State),
		count:    1,
		list:     list.New(),
	}

	for _, word := range wordlist {
		dg.add(word)
	}
	dg.replaceOrRegister(dg.root)

	return dg, nil
}

func (dg *Dawg) add(word string) {
	index, prefix := dg.root.getPrefix(word)

	if hasChildren(prefix) {
		dg.replaceOrRegister(prefix)
	}
	dg.addSuffix(prefix, word[index:])

	dg.length++

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
		dg.list.PushBack(lastchild)
	}
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
func newState(id int) *State {
	return &State{
		id:       id,
		children: Child{},
	}
}
func hasChildren(st *State) bool {
	return st.laststate != nil
}

func lastIndex(st *State) byte {
	return st.lastindex
}

func lastState(st *State) *State {
	return st.laststate
}

func (dg *Dawg) addSuffix(st *State, suffix string) {
	for _, b := range []byte(suffix) {
		dg.count++
		st.children[b] = newState(dg.count)

		st.laststate = st.children[b]
		st.lastindex = b
		st = st.children[b]

	}
	st.final = true
}

func (st *State) getPrefix(word string) (int, *State) {
	var i int
	var b byte
	for _, b = range []byte(word) {
		if st.children[b] == nil {
			return i, st
		}
		st = st.children[b]
		i++
	}
	return i, st
}
