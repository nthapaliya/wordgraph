package dawg

import "sort"

// Dawg is  directed acyclic word graph. Letters are stored as edges of letters,
// and prefixes and matching suffixes are merged. This provides fast lookup, at a
// fraction of the memory consumption of say, a dictionary. It also allows word-related
// operations like looking up words that start with a specific prefix, for example.
//
type Dawg struct {
	root     *State
	register map[string]*State
	count    int
}

// State holds the flags and outgoing edges to other states
//
type State struct {
	final    bool
	children *Child
	hash     string
	id       int
}

// Child is a list of outgoing edges
//
type Child [26]*State

// Root returns a pointer to the root state
//
func (dg Dawg) Root() *State { return dg.root }

// Register returns the register. Required for internal use by library
//
func (dg Dawg) Register() map[string]*State { return dg.register }

// Final returns true if a state is final
//
func (st State) Final() bool { return st.final }

// Children returns the list of children
//
func (st State) Children() *Child { return st.children }

// Hash looks up the states stored hash, which is required for comparision and
// optimization. It does not compute the hash: it merely returns the saved state
// which may or may not be up-to-date
//
func (st State) Hash() string { return st.hash }

// Contains returns true if string exists in dictionary
//
func (dg Dawg) Contains(word string) bool {
	st := dg.root
	for _, b := range []byte(word) {
		b -= offset
		if st.children[b] == nil {
			return false
		}
		st = st.children[b]
	}
	return st.final
}

// List emits all words contained in the dictionary, in order
//
func (dg Dawg) List() []string {
	return dg.ListFrom("")
}

// ListFrom returns all words in the dictionary starting with the given prefix.
// If a prefix does not exist, returns an empty list.
//
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
