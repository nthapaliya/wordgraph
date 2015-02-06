package cdawg_test

import (
	"bytes"
	"io/ioutil"
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

func equals(t *testing.T, a, b cdawg.MDawg) bool {
	if len(a) != len(b) {
		t.Errorf("wanted len %d, got %d", len(a), len(b))
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			t.Errorf("%v != %v", a[i], b[i])
			return false
		}
		for j := range a[i] {
			val1, val2 := a[i][j]&0xfffffdff, b[i][j]&0xfffffdff
			if val1 != val2 {
				t.Errorf("wanted %d, got %d, ", val1, val2)
				t.Errorf("%v != %v", a[i], b[i])
				return false
			}
		}
	}
	return true
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

func testListWordGraph(t *testing.T, cd wordgraph.WordGraph) {
	l := cd.List()
	if len(l) != len(wordlist) {
		t.Errorf("len(wordlist)=%d, len(returnedlist)=%d", len(wordlist), len(l))
		return
	}

	for i := range wordlist {
		if l[i] != wordlist[i] {
			t.Errorf("List(): %s != %s", wordlist[i], l[i])
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

func TestCompressDawg(t *testing.T) {
	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		t.Error(err)
		return
	}
	cd, _ := cdawg.Compress(dg)
	wordSuite(t, cd)
	testListWordGraph(t, cd)
}

func TestMinimizeCDawg(t *testing.T) {
	cd, _ := cdawg.NewFromList(wordlist)
	mm, _ := cd.Minimize()
	wordSuite(t, mm)
	testListWordGraph(t, mm)
}

func TestUnminimize(t *testing.T) {
	cd, _ := cdawg.NewFromList(wordlist)
	mm, _ := cd.Minimize()
	newcd := cdawg.Unminimize(mm)
	if len(cd) != len(newcd) {
		t.Errorf("unminimized cd not same length as original")
		return
	}
	for i := range cd {
		for j := range cd[i] {
			if cd[i][j] != newcd[i][j] {
				t.Errorf("need %d, got %d", cd[i][j], newcd[i][j])
				return
			}
		}
	}
}

func TestReadWriteFromFile(t *testing.T) {
	cd, _ := cdawg.NewFromList(wordlist)
	md, _ := cd.Minimize()
	err := cdawg.MarshalJSON("md.json", md)
	if err != nil {
		t.Error(err)
	}

	md2, err := cdawg.UnmarshalJSON("md.json")
	if err != nil {
		t.Error(err)
	}
	if !equals(t, md, md2) {
		t.Errorf("two are not equal")
	}
}

func TestEncodeDecode(t *testing.T) {
	md, err := cdawg.UnmarshalJSON("../files/md.json")
	if err != nil {
		t.Error(err)
	}
	encoded, _ := cdawg.EncodeToBinary(md)
	decoded, _ := cdawg.DecodeFromBinary(encoded)
	equals(t, md, decoded)
}

func TestBinaryEncodeDecode(t *testing.T) {
	md1, err := cdawg.UnmarshalJSON("../files/md.json")
	if err != nil {
		t.Error(err)
	}
	b1, err := cdawg.EncodeToBinary(md1)
	if err != nil {
		t.Error(err)
		return
	}
	err = ioutil.WriteFile("md.bin.tmp", b1, 0644)
	if err != nil {
		t.Error(err)
		return
	}

	b2, err := ioutil.ReadFile("md.bin.tmp")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(b1, b2) {
		t.Error("bytes not equal!")
	}
	md2, err := cdawg.DecodeFromBinary(b2)
	if err != nil {
		t.Error(err)
		return
	}
	if !equals(t, md1, md2) {
		t.Error("md's not equal")
		return
	}
}

func TestTest(t *testing.T) {
	cdawg.Test()
}
