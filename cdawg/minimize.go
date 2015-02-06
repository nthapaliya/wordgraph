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
	indexBitmask  = (1 << indexShift) - 1
)

// MDawg is a minimized form of CDawg. Arcs are represented as indices to start of
// a list of children, rather than direct access. This compacts the structure as there
// are a lot of zeroes in CDawg. The disadvantage is that looking up words is
// MUCH slower.
//
type MDawg [][]int

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

func isFinal(value int) bool {
	return value&finalBitmask != 0
}

func letter(value int) byte {
	return byte(value & letterBitmask)
}

func firstChild(value int) int {
	return value >> indexShift
}

func next(value int) int {
	if isEOL(value) {
		return 0
	}
	return value + 1
}

func isEOL(value int) bool {
	return value&eolBitmask != 0
}

type tdawg []int

func Test() {
	md, _ := UnmarshalJSON("../files/md.json")
	td := md.create()
	fmt.Println(td[:100])
	i := 1
	eol := false
	for !eol {
		val := td[i]
		l := string(letter(val))
		final := isFinal(val)
		eol = isEOL(val)
		fmt.Printf("val: %d, %s is final? %v, is eol? %v\n", val, l, final, eol)
		i++
	}
}
func (md MDawg) create() tdawg {
	counter := 0
	fromTo := make(map[int]int)
	for i := range md {
		for j := range md[i] {
			v := md[i][j] >> indexShift
			if _, ok := fromTo[v]; !ok {
				fromTo[v] = counter
			}
			counter++
		}
	}
	fmt.Println(len(fromTo), counter)
	td := make(tdawg, counter)

	nextavailable := 0
	for i := range md {
		for j := range md[i] {
			from := md[i][j] >> indexShift
			td[nextavailable] = (md[i][j] & indexBitmask) + (fromTo[from] << indexShift)
			nextavailable++
		}
	}
	return td
}
