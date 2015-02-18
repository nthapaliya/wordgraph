package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/nthapaliya/wordgraph/cdawg"
)

var (
	shortList = []string{"director", "directorate", "directorship", "fellowship"}
)

func main() {
	save(prefixGraph("plane"), "graph.gv")
	// md, _ := cdawg.NewFromList(shortList)
	// save(gvPlain(md), "graph.gv")
}

func prefixGraph(s string) []byte {
	md, _ := cdawg.UnmarshalJSON("../files/cd.json")
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
			buf.WriteString(fmt.Sprintf("%d [label=%c, color=red];\n", v, v&0xff))
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
	buf.WriteString("horizontal=true;\n")

	// Step 7: Finalize
	buf.WriteString("}")
	return buf.Bytes()
}
