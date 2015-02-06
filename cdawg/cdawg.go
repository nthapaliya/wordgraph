package cdawg

import "github.com/nthapaliya/wordgraph/dawg"

const (
	offset        = 'a'
	finalBitmask  = 1 << 8
	eolBitmask    = 1 << 9
	letterBitmask = 0xff
	indexShift    = 10
	indexBitmask  = (1 << indexShift) - 1
)

// CDawg stores arcs as indexes in a 2d matrix instead of pointers to 'floating' structs
// in the heap. Non-editable once made (of course, you can manually change the values, its
// just a matrix, but you'd mess things up)
//
type CDawg [][]int

// NewFromList creates a CDawg from a sorted list of words. If list is unsorted,
// returns nil
//
func NewFromList(wordlist []string) (CDawg, error) {
	dg, err := dawg.NewFromList(wordlist)
	if err != nil {
		return nil, err
	}
	cd, err := Compress(dg)
	if err != nil {
		return nil, err
	}
	return cd, nil
}

// Compress takes the pointer-to-node based Dawg structure and converts it
// into an int matrix CDawg
//
func Compress(dg *dawg.Dawg) (CDawg, error) {
	if ok, err := dg.Verify(); !ok {
		return nil, err
	}
	register := dg.Register()
	cdlen := len(register) + 1

	cd := make(CDawg, cdlen)
	indexOf := make(map[string]int)
	stateAt := make([]*dawg.State, cdlen)

	indexOf[dg.Root().Hash()] = 1
	stateAt[1] = dg.Root()

	for k, v := range register {
		count := 0
		for _, child := range v.Children() {
			if child != nil {
				count++
			}
		}
		if count == 0 {
			indexOf[k] = 0
			stateAt[0] = v
			cd[0] = []int{eolBitmask}
			delete(register, k)
			break
		}
	}
	// up to now, removed final state from register, and put it at index = 0
	// root state is at index = 1

	lastIndex := 1
	currentIndex := 1

	for currentIndex < cdlen {
		currentState := stateAt[currentIndex]
		cd[currentIndex] = []int{}

		for letter, child := range currentState.Children() {
			if child != nil {
				hash := child.Hash()
				if v, ok := register[hash]; ok {
					lastIndex++
					indexOf[hash] = lastIndex
					stateAt[lastIndex] = v
					delete(register, hash)
				}
				val := (letter + offset) + (indexOf[hash] << indexShift)
				if child.Final() {
					val += finalBitmask
				}
				cd[currentIndex] = append(cd[currentIndex], val)
			}
		}
		rowlen := len(cd[currentIndex])
		cd[currentIndex][rowlen-1] += eolBitmask

		currentIndex++
	}
	return cd, nil
}

// Contains returns true if string exists in dictionary
//
func (cd CDawg) Contains(word string) bool {
	index, value := 1, 0
	var ok bool

	for _, b := range []byte(word) {
		if value, ok = hasByteInRow(cd[index], b); !ok {
			return false
		}
		index = value >> indexShift
	}

	return isFinal(value)
}

func hasByteInRow(row []int, b byte) (int, bool) {
	for _, v := range row {
		letter := letter(v)
		if b == letter {
			return v, true
		}
	}
	return 0, false
}
