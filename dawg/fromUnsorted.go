package dawg

// FromUnsorted creates a Dawg from a slice of unsorted data
func FromUnsorted(wordlist []string) *Dawg {
	dg := &Dawg{
		root:     &State{false, &Child{}, ""},
		register: make(map[string]*State),
	}

	for _, word := range wordlist {
		dg.Add(word)
	}
	return dg
}

// Add adds a word to the dictionary
func (dg *Dawg) Add(word string) {

}
