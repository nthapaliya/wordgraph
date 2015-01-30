package dawg

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

func (st *State) addSuffix(suffix string) {
	for _, b := range []byte(suffix) {
		b -= offset
		st.children[b] = &State{false, &Child{}, ""}
		st = st.children[b]
	}
	st.final = true
}

func (st *State) getPrefix(word string) (int, *State) {
	for i, b := range word {
		b -= offset
		if st.children[b] == nil {
			return i, st
		}
		st = st.children[b]
	}
	return 0, st
}
