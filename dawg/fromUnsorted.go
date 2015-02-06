package dawg

// FromUnsorted creates a Dawg from a slice of unsorted data
//
func FromUnsorted(wordlist []string) *Dawg {
	dg := &Dawg{
		root:     &State{final: false, children: &Child{}, hash: "", id: 0},
		register: make(map[string]*State),
		count:    1,
	}

	for _, word := range wordlist {
		dg.Add(word)
	}
	return dg
}

// Add adds a word to the dictionary. Add does not have to be called for words in
// lexicographic order
//
func (dg *Dawg) Add(word string) {
	i, _ := dg.root.getPrefix(word)
	clonedstates := []*State{}
	{
		st := dg.root
		for _, b := range []byte(word[:i]) {
			b -= offset
			clonedstates = append(clonedstates, dg.clone(st.children[b]))
		}
	}
	{
		// prefix clone
		st := clonedstates[len(clonedstates)-1]
		for _, b := range []byte(word[i:]) {
			b -= offset
			dg.count++
			st.children[b] = &State{
				final:    false,
				children: &Child{},
				hash:     "",
				id:       dg.count,
			}
			st = st.children[b]
			clonedstates = append(clonedstates, st)
		}
		st.final = true
	}

	for i := len(clonedstates) - 1; i > 0; i-- {
		dg.register_or_minimize(clonedstates[i])
	}
}

func (dg *Dawg) register_or_minimize(st *State) {
	// hash := st.getHash()
	// if v, ok := dg.register[hash]; !ok {
	//
	// }
}

func (dg *Dawg) clone(st *State) *State {
	dg.count++
	return &State{
		final:    st.final,
		children: st.children,
		hash:     "",
		id:       dg.count,
	}
}
