package wordgraph_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/nthapaliya/wordgraph"
	"github.com/nthapaliya/wordgraph/cdawg"
	"github.com/nthapaliya/wordgraph/dawg"
	"github.com/nthapaliya/wordgraph/trie"
)

var (
	wordlist = wordgraph.LoadFile("files/SWOPODS.txt")
	random   = rand.New(rand.NewSource(time.Now().Unix()))
)

func BenchmarkMDawg(b *testing.B) {
	cd, _ := cdawg.NewFromList(wordlist)
	mn, _ := cd.Minimize()
	benchmarkWordGraph(b, mn)
}

func BenchmarkCDawg(b *testing.B) {
	cd, _ := cdawg.NewFromList(wordlist)
	benchmarkWordGraph(b, cd)
}

func BenchmarkDawg(b *testing.B) {
	cd, _ := dawg.NewFromList(wordlist)
	benchmarkWordGraph(b, cd)
}

func BenchmarkTrie(b *testing.B) {
	cd := trie.NewFromList(wordlist)
	benchmarkWordGraph(b, cd)
}

func BenchmarkMap(b *testing.B) {
	register := make(map[string]bool)
	for _, word := range wordlist {
		register[word] = true
	}
	// contains := func(word string) bool {
	// 	return register[word]
	// }

	for n := 0; n < b.N; n++ {
		word := wordlist[random.Intn(len(wordlist))]
		// contains(word)
		if register[word] {
			// do nothing
		}
	}
}

func benchmarkWordGraph(b *testing.B, cd wordgraph.WordGraph) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		word := wordlist[random.Intn(len(wordlist))]
		cd.Contains(word)
	}
}
