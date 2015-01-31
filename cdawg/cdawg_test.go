package cdawg_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/nthapaliya/wordgraph"
	"github.com/nthapaliya/wordgraph/cdawg"
	"github.com/nthapaliya/wordgraph/dawg"
)

// Notes: now its 77807. nvm

const (
	offset = 97
)

var (
	random   = rand.New(rand.NewSource(time.Now().Unix()))
	wordlist = wordgraph.LoadSame()
)

var badwords = []string{
	"pnkirajc",
	"kebaldadz",
	"dhapasi",
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

func TestCompress(t *testing.T) {
	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		t.Error(err)
		return
	}
	cd := cdawg.Compress(dg)
	wordSuite(t, cd)
}

func TestMinimize(t *testing.T) {
	cd := cdawg.NewFromList(wordlist)
	mm := cd.Minimize()
	wordSuite(t, mm)
}

func TestMinimizeDawg(t *testing.T) {
	dg, _ := dawg.NewFromList(wordlist)
	mm := cdawg.MinimizeDawg(dg)
	wordSuite(t, mm)
}

func TestCDawg(t *testing.T) {
	cd := cdawg.NewFromList(wordlist)
	for i := range cd {
		i++
	}
}
func TestList(t *testing.T) {
	cd := cdawg.NewFromList(wordlist)
	l := cd.List()
	if len(l) != len(wordlist) {
		t.Errorf("Error. len(wordlist)=%d, len(returnedlist)=%d", len(wordlist), len(l))
		return
	}

	for i := range wordlist {
		if l[i] != wordlist[i] {
			t.Errorf("List(): results don't match")
		}
	}

	l = cd.ListFrom("applx")
	if len(l) != 0 {
		t.Errorf("testing applx: len(l)=%d, expected 0", len(l))
	}

	l = cd.ListFrom("apply")
	if len(l) != 2 {
		t.Errorf("testing apply: len(l)=%d, expected 2", len(l))
	}
}
