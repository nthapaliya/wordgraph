package dawg

import (
	"fmt"
	"hash/fnv"
	"io"
)

func (st *State) getHash() string {
	if st.hash == "" {
		hash := make([]string, 27)
		if st.final {
			hash[0] = "final"
		}
		for i, c := range st.children {
			if c != nil {
				hash[i+1] = c.getHash()
			} else {
				hash[i+1] = "nil"
			}
		}
		st.hash = hashFNV64a(fmt.Sprintf("%v", hash))
	}
	return st.hash
}

func hashFNV64a(in string) string {
	h := fnv.New64a()
	io.WriteString(h, in)

	return fmt.Sprintf("%x", h.Sum(nil))
}
