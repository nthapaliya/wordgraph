package cdawg

// Contains returns true if string exists in dictionary
//
func (cd CDawg) Contains(word string) bool {
	index := 0
	val := 0
	for _, b := range []byte(word) {
		b -= offset
		val = cd[index][b]
		index = val >> 1
		if index == 0 {
			return false
		}
	}
	return val&1 == 1
}

// Contains returns true if string exists in dictionary
//
func (cd MDawg) Contains(word string) bool {
	var state, index, value int
	var ok bool

	for _, b := range []byte(word) {
		if value, ok = hasByteInRow(cd[index], b); !ok {
			return false
		}
		index = value >> indexShift
		state = value
	}

	return decode(state).final
}

func hasByteInRow(row []int, b byte) (int, bool) {
	for _, v := range row {
		ss := decode(v)
		if b == ss.letter {
			return v, true
		}
	}
	return 0, false
}
