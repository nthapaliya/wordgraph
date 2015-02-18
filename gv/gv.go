package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/nthapaliya/wordgraph/cdawg"
)

var (
	cdAll, _   = cdawg.UnmarshalJSON("../files/cd.json")
	cdSmall, _ = cdawg.UnmarshalJSON("../files/cd.short.json")
)

func main() {
	// save(gvPlain(cdSmall), "graph.gv")
	save(prefixGraph("plane"), "graph.gv")
}

func prefixGraph(s string) []byte {
	md := cdAll
	listNew := md.ListFrom(s)
	md, _ = cdawg.NewFromList(listNew)

	return gvPlain(md)

}

func save(b []byte, filename string) error {
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func gvPlain(md cdawg.CDawg) []byte {
	// Step 1: Create buffer and write header
	var buf bytes.Buffer
	buf.WriteString("digraph Dawg {\n")

	// Step 2: Enumerate all nodes
	nodes := []int{}
	for i := 1; i < len(md); i++ {
		nodes = append(nodes, md[i]...)
	}

	// Step 3: Write nodes and labels
	for _, v := range nodes {
		if v&(1<<8) != 0 {
			buf.WriteString(fmt.Sprintf("%d [label=%c, color=red, peripheries=2];\n", v, v&0xff))
		} else {
			buf.WriteString(fmt.Sprintf("%d [label=%c];\n", v, v&0xff))
		}
	}

	// Step 4: Capture all edges
	for _, src := range nodes {
		childIndex := src >> 10
		for _, dst := range md[childIndex] {
			if dst != 512 {
				buf.WriteString(fmt.Sprintf("%d->%d;\n", src, dst))
			}
		}
	}

	// Step 5: Other attributes
	buf.WriteString("concentrate=true;\n")
	buf.WriteString("mclimit=10.0;\n")
	buf.WriteString("rankdir=LR;\n")

	// Step 7: Finalize
	buf.WriteString("}")
	return buf.Bytes()
}
