package cdawg

// Contains returns true if string exists in dictionary
//
func (cd CDawg) Contains(word string) bool {
	index := 1
	val := 0
	for _, b := range []byte(word) {
		b -= offset
		val = cd[index][b]
		index = val >> 1
		if val == 0 {
			return false
		}
	}
	return val&1 == 1
}

// Contains returns true if string exists in dictionary
//
func (md MDawg) Contains(word string) bool {
	var state, index, value int
	var ok bool

	index = 1
	for _, b := range []byte(word) {
		if value, ok = hasByteInRow(md[index], b); !ok {
			return false
		}
		index = value >> indexShift
		state = value
	}

	return isFinal(state)
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
