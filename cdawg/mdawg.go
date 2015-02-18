package cdawg

import (
	"errors"

	"github.com/nthapaliya/wordgraph/dawg"
)

// MDawg is a minimized form of CDawg. All arcs and states are placed in one flat array.
// Childnodes are denoted in pretty much the exact same way, but with an EOL (end of list)
// flag to denote the end of a list of children
//
type MDawg []int

func isFinal(value int) bool {
	return value&finalBitmask != 0
}

func letter(value int) byte {
	return byte(value & letterBitmask)
}

func firstChild(value int) int {
	return value >> indexShift
}

func isEOL(value int) bool {
	return value&eolBitmask != 0
}

////////////////////////////////////////////////////////////////////////////////

// MinimizeDawg minimizes a dawg.Dawg. This may perform better on your system
// than a CDawg.
//
func MinimizeDawg(dg *dawg.Dawg) (MDawg, error) {
	cd, err := Compress(dg)
	if err != nil {
		return nil, err
	}
	return cd.Minimize()
}

// Minimize creates a flat array MDawg from the given CDawg structure
// TODO: add more safety checks
//
func (cd CDawg) Minimize() (MDawg, error) {
	if len(cd) == 0 {
		return nil, errors.New("Empty CDawg passed in for minimization")
	}
	// Step 1: Map from old [i][j]->[i'], where i' = sum of j_0..j_n-1
	counter := 0
	newIndex := make(map[int]int)
	for i := range cd {
		newIndex[i] = counter
		counter += len(cd[i])
	}
	td := make(MDawg, counter)

	nextavailable := 0
	for i := range cd {
		for j := range cd[i] {
			oldIndex := firstChild(cd[i][j])
			td[nextavailable] = (cd[i][j] & indexBitmask) + (newIndex[oldIndex] << indexShift)
			nextavailable++
		}
	}
	return td, nil
}

////////////////////////////////////////////////////////////////////////////////

// Contains returns true if string exists in dictionary
//
func (md MDawg) Contains(word string) bool {
	index, val := 1, md[1]
	var ok bool
	for _, b := range []byte(word) {
		if val, ok = hasByteBeforeEOL(md[index:], b); !ok {
			return false
		}
		index = firstChild(val)
	}
	return val&finalBitmask != 0
}

func hasByteBeforeEOL(values []int, b byte) (int, bool) {
	for _, val := range values {
		if letter(val) == b {
			return val, true
		} else if isEOL(val) {
			return 0, false
		}
	}
	return 0, false
}

////////////////////////////////////////////////////////////////////////////////

// List returns a sorted list of all words contained in dictionary
//
func (md MDawg) List() []string {
	return md.ListFrom("")
}

// ListFrom returns a list of words that start with given prefix. If prefix doesn't
// exist in dictionary, returns an empty list
//
func (md MDawg) ListFrom(prefix string) []string {
	value := 1 << indexShift // so that letter(value)==0. we'll be adding it after the fact.
	{
		index := 1
		var ok bool

		for _, b := range []byte(prefix) {
			if value, ok = hasByteBeforeEOL(md[index:], b); !ok {
				return []string{}
			}
			index = firstChild(value)
		}
	}
	// if we exit from above, index will now have our the last good state in the prefix
	f := md.traverseMDawg
	return readFromStream(f, value, prefix)
}

func (md MDawg) traverseMDawg(val int, prefix []byte, stream chan string) {
	if val == eolBitmask {
		return
	}
	if isFinal(val) {
		stream <- string(prefix)
	}
	for _, value := range md[firstChild(val):] {
		md.traverseMDawg(value, append(prefix, letter(value)), stream)
		if isEOL(value) { // we have to stop at the EOL
			return
		}
	}
}
