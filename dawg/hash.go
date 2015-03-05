package dawg

import (
	"fmt"
	"hash/fnv"
	"io"
	"strings"
)

var runeLen = 128

func (st *State) getHash() string {
	if st.hash == "" {
		hash := []string{}

		if st.final {
			hash = append(hash, "final")
		}

		for i := 0; i < runeLen; i++ {
			if c, ok := st.children[byte(i)]; ok {
				hash = append(hash, fmt.Sprintf("(%d->%s)", i, c.getHash()))
			}
		}

		strepr := strings.Join(hash, " ")
		st.hash = hashFNV64a(strepr)
	}
	return st.hash
}

func hashFNV64a(in string) string {
	h := fnv.New64a()
	io.WriteString(h, in)

	return fmt.Sprintf("%x", h.Sum(nil))
}
