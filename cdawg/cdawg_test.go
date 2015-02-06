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
	wordlist = wordgraph.LoadFile("../files/SWOPODS.txt")
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

func equals(t *testing.T, a, b cdawg.CDawg) bool {
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

func testContains(t *testing.T, cd wordgraph.WordGraph) {
	var good, bad int
	for _, word := range wordlist {
		if !cd.Contains(word) {
			bad++
		} else {
			good++
		}
	}
	if bad != 0 {
		t.Errorf("Contains(word): bad = %d, expected 0", bad)
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
		t.Errorf("Contains(word): good = %d, expected 0", good)
	}
}

func testList(t *testing.T, cd wordgraph.WordGraph) {
	l := cd.List()
	if len(l) != len(wordlist) {
		t.Errorf("len(wordlist)=%d, len(returnedlist)=%d", len(wordlist), len(l))
		return
	}

	for i := range wordlist {
		if l[i] != wordlist[i] {
			t.Fatalf("List(): %s != %s", wordlist[i], l[i])
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
	if testing.Short() {
		t.Skip("skip creating and compressing dawg in short mode.")
	}

	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		t.Error(err)
		return
	}
	cd, _ := cdawg.Compress(dg)
	testContains(t, cd)
	testList(t, cd)
}

func TestCDawgUnmarshal(t *testing.T) {
	cd, _ := cdawg.UnmarshalJSON("../files/cd.json")
	testContains(t, cd)
	testList(t, cd)
}

func TestReadWriteFromFile(t *testing.T) {
	cd1, _ := cdawg.NewFromList(wordlist)
	err := cdawg.MarshalJSON("/tmp/cd.json", cd1)
	if err != nil {
		t.Error(err)
	}

	cd2, err := cdawg.UnmarshalJSON("/tmp/cd.json")
	if err != nil {
		t.Error(err)
	}
	if !equals(t, cd1, cd2) {
		t.Errorf("two are not equal")
	}
}

func TestEncodeDecode(t *testing.T) {
	cd, err := cdawg.UnmarshalJSON("../files/cd.json")
	if err != nil {
		t.Error(err)
	}
	encoded, _ := cdawg.EncodeToBytes(cd)
	decoded, _ := cdawg.DecodeFromBytes(encoded)
	equals(t, cd, decoded)
}

func TestBinaryEncodeDecode(t *testing.T) {
	cd1, err := cdawg.UnmarshalJSON("../files/cd.json")
	if err != nil {
		t.Error(err)
	}
	b1, err := cdawg.EncodeToBytes(cd1)
	if err != nil {
		t.Error(err)
		return
	}
	err = ioutil.WriteFile("/tmp/cd.bin.tmp", b1, 0644)
	if err != nil {
		t.Error(err)
		return
	}

	b2, err := ioutil.ReadFile("/tmp/cd.bin.tmp")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(b1, b2) {
		t.Error("bytes not equal!")
	}
	cd2, err := cdawg.DecodeFromBytes(b2)
	if err != nil {
		t.Error(err)
		return
	}
	if !equals(t, cd1, cd2) {
		t.Error("cd's not equal")
		return
	}
}

func TestMinimizeCDawg(t *testing.T) {
	cd, err := cdawg.UnmarshalJSON("../files/cd.json")
	mm, err := cd.Minimize()
	if err != nil {
		t.Fatal(err)
	}
	testContains(t, mm)
}
