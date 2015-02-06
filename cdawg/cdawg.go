package cdawg

import "github.com/nthapaliya/wordgraph/dawg"

const offset = 97

// CDawg stores arcs as indexes in a 2d matrix instead of pointers
//
type CDawg [][]int

// NewFromList creates a CDawg from a sorted list of words. If list is unsorted,
// returns nil
//
func NewFromList(wordlist []string) (CDawg, error) {
	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		return nil, err
	}
	cd, err := Compress(dg)
	if err != nil {
		return nil, err
	}
	return cd, nil
}
