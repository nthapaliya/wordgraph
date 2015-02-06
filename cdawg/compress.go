package cdawg

import "github.com/nthapaliya/wordgraph/dawg"

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
			cd[0] = make([]int, 26)
			delete(register, k)
			break
		}
	}
	// up to now, removed final state and put it as index = 0
	// root state is index = 1
	lastIndex := 1
	currentIndex := 1

	for len(register) != 0 || currentIndex < cdlen {
		currentState := stateAt[currentIndex]
		cd[currentIndex] = make([]int, 26)

		for letter, child := range currentState.Children() {
			if child != nil {
				hash := child.Hash()
				if v, ok := register[hash]; ok {
					lastIndex++
					indexOf[hash] = lastIndex
					stateAt[lastIndex] = v
					delete(register, hash)
				}
				cd[currentIndex][letter] = indexOf[hash] << 1
				if child.Final() {
					cd[currentIndex][letter]++
				}
			}

		}
		currentIndex++
	}
	return cd, nil
}
