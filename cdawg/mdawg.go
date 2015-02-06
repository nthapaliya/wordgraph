package cdawg

import (
	"errors"

	"github.com/nthapaliya/wordgraph/dawg"
)

// MDawg is a minimized form of CDawg. All arcs and states are placed in one flat array.
// Childnodes are denoted in pretty much the exact same way, but with an EOL (end of list)
// flag to denote the end of a list of children
//
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

// MinimizeDawg minimizes a dawg.Dawg. Not recommended for actual use. Use CDawg instead.
// This is more of a curiosity
//
func MinimizeDawg(dg *dawg.Dawg) (MDawg, error) {
	cd, err := Compress(dg)
	if err != nil {
		return nil, err
	}
	return cd.Minimize()
}

// Minimize creates a flat array MDawg from the given CDawg structure
//
func (cd CDawg) Minimize() (MDawg, error) {
	if len(cd) == 0 {
		return nil, errors.New("Empty Compressed Dawg passed in for minimization")
	}
	counter := 0
	newIndex := make(map[int]int)
	for i := range cd {
		for range cd[i] {
			if _, ok := newIndex[i]; !ok {
				newIndex[i] = counter
			}
			counter++
		}
	}
	td := make(MDawg, counter)

	nextavailable := 0
	for i := range cd {
		for j := range cd[i] {
			oldIndex := cd[i][j] >> indexShift
			td[nextavailable] = (cd[i][j] & indexBitmask) + (newIndex[oldIndex] << indexShift)
			nextavailable++
		}
	}
	return td, nil
}

// Contains returns true if string exists in dictionary
//
func (md MDawg) Contains(word string) bool {
	index, val := 1, md[1]
	var ok bool
	for _, b := range []byte(word) {
		if val, ok = hasByteBeforeEOL(md[index:], b); !ok {
			return false
		}
		index = val >> indexShift
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
