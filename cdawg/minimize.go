package cdawg

import (
	"fmt"

	"github.com/nthapaliya/wordgraph/dawg"
)

const (
	finalBitmask  = 1 << 8
	letterBitmask = 0xff
	indexShift    = 10
)

// MDawg is a minimized form of CDawg. Arcs are represented as indices to start of
// a list of children, rather than direct access. This compacts the structure as there
// are a lot of zeroes in CDawg. The disadvantage is that looking up words is
// MUCH slower.
//
type MDawg [][]int

type smallState struct {
	letter byte
	index  int
	final  bool
}

// MinimizeDawg minimizes a dawg.Dawg
func MinimizeDawg(dg *dawg.Dawg) MDawg {
	cd := Compress(dg)
	return cd.Minimize()
}

// Minimize minimises a CDawg
func (cd CDawg) Minimize() MDawg {
	matrix := make([][]int, len(cd))
	for i, row := range cd {
		for letter, val := range row {
			if val != 0 {
				final := val&1 == 1
				index := val >> 1
				matrix[i] = append(matrix[i], encode(letter, index, final))
			}
		}
	}
	return matrix
}

func encode(letter, index int, final bool) int {
	endval := letter + offset
	if final {
		endval |= finalBitmask
	}
	endval += index << indexShift
	return int(endval)
}

func decode(value int) *smallState {
	letter := byte(value & letterBitmask)
	index := value >> indexShift
	final := false
	if (value & finalBitmask) != 0 {
		final = true
	}
	return &smallState{
		letter: letter,
		index:  int(index),
		final:  final,
	}
}

func (ss smallState) String() string {
	return fmt.Sprintf("%c, %d, %v\n", ss.letter, ss.index, ss.final)
}
