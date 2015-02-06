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

func BenchmarkCDawg(b *testing.B) {
	cd, err := cdawg.UnmarshalJSON("./files/cd.json")
	if err != nil {
		b.Fatal(err)
	}
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

func benchmarkWordGraph(b *testing.B, wg wordgraph.WordGraph) {
	shuffled := wordgraph.Shuffle(wordlist)
	length := len(shuffled)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		wg.Contains(shuffled[n%length])
	}
}
