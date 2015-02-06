package cdawg

// Unminimize is the opposite of Minimize. Takes a MDawg returns an 'expanded' version
// that performs faster.
//
func Unminimize(md MDawg) (cd CDawg) {
	cd = make([][]int, len(md))
	cd[0] = make([]int, 26)

	for i := 1; i < len(md); i++ {
		cd[i] = make([]int, 26)
		for _, val := range md[i] {
			letter := (val & letterBitmask) - offset
			cd[i][letter] = (val >> indexShift) << 1
			if val&finalBitmask != 0 {
				cd[i][letter]++
			}
		}
	}
	return cd
}
