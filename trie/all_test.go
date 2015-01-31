package trie_test

import (
	"testing"

	"github.com/nthapaliya/wordgraph"
	"github.com/nthapaliya/wordgraph/trie"
)

// running wc -l on the file
const (
	wc = 267751
)

var wordlist = wordgraph.LoadSame()

func TestNew(t *testing.T) {
	tr := trie.NewFromList(wordlist)
	for _, word := range wordlist {
		if !tr.Contains(word) {
			t.Errorf("error!")
		}
	}
}

func TestListAll(t *testing.T) {
	tr := trie.NewFromList(wordlist)
	l := tr.List()

	if len(l) != len(wordlist) {
		t.Errorf("retrieved list different from added list")
	}

	falsecount := 0
	for i, word := range wordlist {
		if l[i] != word {
			falsecount++
		}
	}
	if falsecount != 0 {
		t.Errorf("list len equal but produces different words")
	}

	l = tr.ListFrom("applx")
	if len(l) != 0 {
		t.Errorf("producing list when it shouldn't be")
	}
	l = tr.ListFrom("apply")
	if len(l) != 2 {
		t.Errorf("investigate here, wrong output")
	}
}
