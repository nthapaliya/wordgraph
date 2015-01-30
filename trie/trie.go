package trie

// Basic constants
const offset = 97

// New returns pointer to an emptiy Trie, with initialized root
func New() (tr *Trie) {
	return &Trie{
		root: &state{
			final:    false,
			children: &child{},
			hash:     "",
		},
	}
}

// NewFromList creates a trie from a list of strings (sorted or unsorted)
//
func NewFromList(list []string) (tr *Trie) {
	tr = New()

	for _, word := range list {
		tr.Add(word)
	}
	return tr
}

// Add adds an individual string to the dictionary
//
func (tr *Trie) Add(word string) {
	r := tr.root

	for _, b := range []byte(word) {
		if r.children[b-offset] == nil {
			r.children[b-offset] = &state{final: false, children: &child{}}
		}
		r = r.children[b-offset]
	}
	r.final = true
}
