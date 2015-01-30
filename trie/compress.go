package trie

import "github.com/nthapaliya/wordgraph/cdawg"

// Compress ...
func (tr Trie) Compress() cdawg.CDawg {
	if tr.root.hash != "" {
		return nil
	}

	return nil
}
