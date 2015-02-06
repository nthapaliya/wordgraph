package wordgraph

// WordGraph ...
type WordGraph interface {
	Contains(string) bool
	// Add(string)
	List() []string
	ListFrom(prefix string) []string
}
