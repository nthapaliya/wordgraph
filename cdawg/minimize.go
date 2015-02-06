package cdawg

import (
	"fmt"

	"github.com/nthapaliya/wordgraph/dawg"
)

const (
	finalBitmask  = 1 << 8
	eolBitmask    = 1 << 9
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
func MinimizeDawg(dg *dawg.Dawg) (MDawg, error) {
	cd, err := Compress(dg)
	if err != nil {
		return nil, err
	}
	return cd.Minimize()
}

// Minimize minimises a CDawg
func (cd CDawg) Minimize() (MDawg, error) {
	matrix := make([][]int, len(cd))
	// null state
	matrix[0] = []int{eolBitmask}

	for i := 1; i < len(cd); i++ {
		row := cd[i]
		for letter, val := range row {
			if val != 0 {
				encodedvalue := letter + offset
				// If final, append
				if val&1 == 1 {
					encodedvalue += finalBitmask
				}
				// val >> 1 == index
				// add index << indexShift
				encodedvalue += (val >> 1) << indexShift
				matrix[i] = append(matrix[i], encodedvalue)
			}
		}
		rowlen := len(matrix[i])
		matrix[i][rowlen-1] |= eolBitmask
	}
	return matrix, nil
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
