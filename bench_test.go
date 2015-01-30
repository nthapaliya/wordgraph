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

func BenchmarkMin(b *testing.B) {
	cd := cdawg.NewFromList(wordlist)
	mn := cd.Minimize()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		word := wordlist[random.Intn(len(wordlist))]
		mn.Contains(word)
	}
}

func BenchmarkCDawg(b *testing.B) {
	cd := cdawg.NewFromList(wordlist)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		word := wordlist[random.Intn(len(wordlist))]
		cd.Contains(word)
	}
}

func BenchmarkDawg(b *testing.B) {
	cd, _ := dawg.NewFromList(wordlist)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		word := wordlist[random.Intn(len(wordlist))]
		cd.Contains(word)
	}
}

func BenchmarkTrie(b *testing.B) {
	cd := trie.NewFromList(wordlist)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		word := wordlist[random.Intn(len(wordlist))]
		cd.Contains(word)
	}
}

func BenchmarkMap(b *testing.B) {
	register := make(map[string]bool)
	for _, word := range wordlist {
		register[word] = true
	}
	// contains := func(word string) bool {
	// 	return register[word]
	// }

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		word := wordlist[random.Intn(len(wordlist))]
		// contains(word)
		if register[word] {
			// do nothing
		}
	}
}