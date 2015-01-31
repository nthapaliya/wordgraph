package cdawg

import "github.com/nthapaliya/wordgraph/dawg"

// Compress takes the pointer-to-node based Dawg structure and converts it
// into an int matrix CDawg
//
func Compress(dg *dawg.Dawg) CDawg {

	indices := getIndices(dg)
	reg := populateReg(indices)
	matrix := make([][]int, len(indices))

	for i := range matrix {
		matrix[i] = make([]int, 26)

		for j, c := range indices[i].Children() {
			if c != nil {
				if c.Final() {
					matrix[i][j] = 1
				}
				matrix[i][j] += reg[c.Hash()] << 1
			}
		}
	}
	return matrix
}

func getIndices(dg *dawg.Dawg) []*dawg.State {
	indices := []*dawg.State{}
	encountered := make(map[string]bool)
	q := newQueue(100)
	q.push(dg.Root())

	for q.count != 0 {
		st := q.pop()
		if !encountered[st.Hash()] {
			encountered[st.Hash()] = true
			indices = append(indices, st)

			for _, c := range st.Children() {
				if c != nil {
					q.push(c)
				}
			}
		}

	}
	return indices
}

func populateReg(indices []*dawg.State) map[string]int {
	reg := make(map[string]int)
	for i, st := range indices {
		reg[st.Hash()] = i
	}
	return reg
}
