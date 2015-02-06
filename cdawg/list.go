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
	value := 1 << indexShift
	{
		index := 1
		var ok bool

		for _, b := range []byte(prefix) {
			if value, ok = hasByteInRow(cd[index], b); !ok {
				return []string{}
			}
			index = value >> indexShift
		}
	}
	f := cd.traverseMDawg
	return readFromStream(f, value, prefix)
}

func (cd CDawg) traverseMDawg(val int, prefix []byte, stream chan string) {
	if val == eolBitmask {
		return
	}
	if val&finalBitmask != 0 {
		stream <- string(prefix)
	}
	for _, value := range cd[val>>indexShift] {
		cd.traverseMDawg(value, append(prefix, offset+byte(value&letterBitmask)), stream)
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

func (md MDawg) List() []string {
	return nil
}

func (md MDawg) ListFrom(word string) []string {
	return nil
}
