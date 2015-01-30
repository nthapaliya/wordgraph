package trie

import (
	"fmt"
	"hash/fnv"
	"io"
	"strings"
)

func (st *state) getHash() string {
	if st.hash != "" {
		return st.hash
	}

	outbuffer := []string{}

	if st.final {
		outbuffer = append(outbuffer, "final")
	}
	for _, v := range st.children {
		if v != nil {
			outbuffer = append(outbuffer, v.getHash())
		} else {
			outbuffer = append(outbuffer, "nil")
		}
	}
	st.hash = hashFNV32a(strings.Join(outbuffer, ","))
	return st.hash

}

func hashFNV64a(in string) string {
	h := fnv.New64a()
	io.WriteString(h, in)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashFNV32a(in string) string {
	h := fnv.New32a()
	io.WriteString(h, in)

	return fmt.Sprintf("%x", h.Sum(nil))
}
