package wordgraph_test

import (
	"sort"
	"testing"

	"github.com/nthapaliya/wordgraph"
)

func TestShuffle(t *testing.T) {
	wordlist := wordgraph.LoadFile("files/SWOPODS.txt")
	l := wordgraph.Shuffle(wordlist)
	if len(l) != len(wordlist) {
		t.Errorf("lengths dont match")
	}
	sort.Strings(l)
	for i := range l {
		if l[i] != wordlist[i] {
			t.Errorf("sorted lists not equal")
		}
	}
}
