package dawg_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/nthapaliya/wordgraph"
	"github.com/nthapaliya/wordgraph/dawg"
)

// Notes: now its 77807. nvm

var (
	random   = rand.New(rand.NewSource(time.Now().Unix()))
	wordlist = wordgraph.LoadFile("../files/SWOPODS.txt")
)

var badwords = []string{
	"thisisa",
	"aaplesnii",
	"paanipary",
	"meronaam",
	"tveherau",
	"asdfasfddafasfd",
	"stdsawerjajaja",
	"appleappl",
	"testif",
	"verificatio",
}

func wordSuite(t *testing.T, cd wordgraph.WordGraph) {
	var good, bad int
	for _, word := range wordlist {
		if !cd.Contains(word) {
			bad++
		} else {
			good++
		}
	}
	if bad != 0 {
		t.Errorf("bad = %d", bad)
	}
	good, bad = 0, 0
	for _, word := range badwords {
		if !cd.Contains(word) {
			bad++
		} else {
			good++
		}
	}
	if good != 0 {
		t.Errorf("good = %d", good)
	}
}

func TestVerify(t *testing.T) {
	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		t.Error(err)
	}
	if err := dg.Verify(); err != nil {
		t.Error(err)
	}
}

func TestExists(t *testing.T) {
	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		t.Error(err)
		return
	}
	wordSuite(t, dg)
}

func TestList(t *testing.T) {
	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		t.Fatal(err)
	}
	l := dg.List()
	if expected, got := len(wordlist), len(l); expected != got {
		t.Fatalf("lens don't match, expected %d, got %d", expected, got)
	}

	for i := range wordlist {
		if expected, got := wordlist[i], l[i]; expected != got {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	}
	l = dg.ListFrom("applx")
	if len(l) != 0 {
		t.Fatal("returning items from prefix that doesn't exist")
	}
	l = dg.ListFrom("apply")
	if len(l) != 2 {
		t.Fatal("returned list seems to be wrong, investigate here")
	}
}
