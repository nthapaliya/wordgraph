package wordgraph

// WordGraph ...
type WordGraph interface {
	Contains(string) bool
	List() []string
}
