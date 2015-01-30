package cdawg

// Contains returns true if string exists in dictionary
//
func (cd MDawg) Contains(word string) bool {
	state := 0
	index := 0
	value := 0
	ok := false

	for _, b := range []byte(word) {
		if value, ok = cd.has(index, b); !ok {
			return false
		}
		index = value >> indexShift
		state = value
	}

	return decode(state).final
}

func (cd MDawg) has(state int, b byte) (int, bool) {
	for _, v := range cd[state] {
		ss := decode(v)
		if b == ss.letter {
			return v, true
		}
	}
	return 0, false
}

// Contains returns true if string exists in dictionary
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
