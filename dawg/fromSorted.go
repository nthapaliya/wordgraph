package dawg

import "sort"

const offset = 'a'

// NewFromList creates a new Dawg from a slice of sorted strings
//
func NewFromList(wordlist []string) (*Dawg, error) {
	if !sort.StringsAreSorted(wordlist) {
		return FromUnsorted(wordlist)
	}

	dg := &Dawg{
		root:     &State{false, &Child{}, "", 0},
		register: make(map[string]*State),
		count:    1,
	}

	for _, word := range wordlist {
		index, prefix := dg.root.getPrefix(word)

		if hasChildren(prefix) {
			dg.replaceOrRegister(prefix)
		}
		dg.addSuffix(prefix, word[index:])
	}
	dg.replaceOrRegister(dg.root)

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

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func hasChildren(st *State) bool {
	for _, v := range st.children {
		if v != nil {
			return true
		}
	}
	return false
}

func lastIndex(st *State) int {
	_, i, _ := last(st)
	return i
}

func lastState(st *State) *State {
	s, _, _ := last(st)
	return s
}

func last(st *State) (*State, int, error) {
	var last *State
	var indx int

	for i, v := range st.children {
		if v != nil {
			last = v
			indx = i
		}
	}
	return last, indx, nil
}

func (dg *Dawg) addSuffix(st *State, suffix string) {
	for _, b := range []byte(suffix) {
		b -= offset
		st.children[b] = &State{false, &Child{}, "", dg.count + 1}
		st = st.children[b]
		dg.count++
	}
	st.final = true
}

func (st *State) getPrefix(word string) (int, *State) {
	var i int
	var b byte
	for _, b = range []byte(word) {
		b -= offset
		if st.children[b] == nil {
			return i, st
		}
		st = st.children[b]
		i++
	}
	return i, st
}
