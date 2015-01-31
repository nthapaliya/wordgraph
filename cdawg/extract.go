package cdawg

import "sort"

// List returns a sorted list of all words contained in dictionary
//
func (cd CDawg) List() []string {
	return cd.ListFrom("")
}

// ListFrom returns a list of words that start with given prefix. If prefix doesn't
// exist in dictionary, returns an empty list
//
func (cd CDawg) ListFrom(prefix string) []string {
	val := 0
	{
		index := 0
		for _, b := range []byte(prefix) {
			b -= offset
			val = cd[index][b]
			index = val >> 1
			if val == 0 {
				return []string{}
			}
		}
	}

	stream := make(chan string, 1000)
	go func() {
		cd.traverseCdawg(val, []byte(prefix), stream)
		close(stream)
	}()

	outputlist := []string{}
	for word := range stream {
		outputlist = append(outputlist, word)
	}
	sort.Strings(outputlist)
	return outputlist
}

func (cd CDawg) traverseCdawg(val int, prefix []byte, stream chan string) {
	if val&1 == 1 {
		stream <- string(prefix)
	}

	for i, val := range cd[val>>1] {
		if val != 0 {
			cd.traverseCdawg(val, append(prefix, byte(i+offset)), stream)
		}
	}
}

// ListFrom returns a list of words that start with given prefix. If prefix doesn't
// exist in dictionary, returns an empty list
//
func (md MDawg) ListFrom(prefix string) []string {
	return nil
}
