package dawg

// Contains returns true if string exists in dictionary
func (dg Dawg) Contains(word string) bool {
	st := dg.root
	for _, b := range []byte(word) {
		b -= offset
		if st.children[b] == nil {
			return false
		}
		st = st.children[b]
	}
	return st.final
}
