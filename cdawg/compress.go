package cdawg

import "github.com/nthapaliya/wordgraph/dawg"

// Compress takes the pointer-to-node based Dawg structure and converts it
// into an int matrix CDawg
func Compress(dg *dawg.Dawg) CDawg {
	matrix := make([][]int, len(dg.Register())+1)
	for i := range matrix {
		matrix[i] = make([]int, 26)
	}
	r := make(map[string]int)

	indices := make([]*dawg.State, len(dg.Register())+1)
	indices[0] = dg.Root()

	index := 1
	for k, v := range dg.Register() {
		_, ok := r[k]
		if !ok {
			r[k] = index
			indices[index] = v
			index++
		}
	}
	// all set up

	for i, st := range dg.Root().Children() {
		matrix[0][i] = r[st.Hash()] << 1
		if st.Final() {
			matrix[0][i]++
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j, c := range indices[i].Children() {
			if c != nil {
				matrix[i][j] = r[c.Hash()] << 1
				if c.Final() {
					matrix[i][j]++
				}
			}
		}
	}
	return matrix
}
