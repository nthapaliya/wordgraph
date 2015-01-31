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

	f := cd.traverseCdawg
	return readFromStream(f, val, prefix)
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

// List returns a sorted list of all words contained in dictionary
//
func (md MDawg) List() []string {
	return md.ListFrom("")
}

// ListFrom returns a list of words that start with given prefix. If prefix doesn't
// exist in dictionary, returns an empty list
//
func (md MDawg) ListFrom(prefix string) []string {
	var value int
	{
		var index int
		var ok bool

		for _, b := range []byte(prefix) {
			if value, ok = hasByteInRow(md[index], b); !ok {
				return []string{}
			}
			index = value >> indexShift
			// state = value
		}
	}
	f := md.traverseMDawg
	return readFromStream(f, value, prefix)
}

func (md MDawg) traverseMDawg(val int, prefix []byte, stream chan string) {
	if val&finalBitmask != 0 {
		stream <- string(prefix)
	}
	for _, val := range md[val>>indexShift] {
		md.traverseMDawg(val, append(prefix, byte(val&letterBitmask)), stream)
	}

}

func readFromStream(f func(int, []byte, chan string), val int, prefix string) []string {
	stream := make(chan string, 1000)
	go func() {
		f(val, []byte(prefix), stream)
		close(stream)
	}()

	outputlist := []string{}
	for word := range stream {
		outputlist = append(outputlist, word)
	}
	sort.Strings(outputlist)
	return outputlist

}
