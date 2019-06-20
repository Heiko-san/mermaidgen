package flowchart_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/flowchart"
)

// Working with Flowcharts
func ExampleFlowchart() {
	f := flowchart.NewFlowchart()
	// define some Subgraphs
	// see Subgraph for details about what you can do with them
	sg1 := f.AddSubgraph("sg1")
	sg1.Title = "Generating Flowchart graphs"
	// define some Nodes
	// see Node and NodeStyle for details about what you can do with them
	n1 := sg1.AddNode("n1")
	n2 := sg1.AddNode("n2")
	// define some Edges
	// see Edge and EdgeStyle for details about what you can do with them
	f.AddEdge(n1, n2)
	// you can lookup already defined entities
	// but if the ID is not found nil is returned
	sgX := f.GetSubgraph("sg2")
	nX := f.GetNode("n3")
	eX := f.GetEdge(2)
	fmt.Println(sgX, nX, eX)
	// same way you get nil if you try to add an existing ID
	sgY := f.AddSubgraph("sg1")
	nY := f.AddNode("n1")
	// actually this works, since Edges don't have a real ID
	eY := f.AddEdge(n1, n2)
	fmt.Println(sgY, nY, eY)
	// you can iterate over all entities defined in the Flowcharts
	// and all of its Subgraphs
	for _, item := range f.ListSubgraphs() {
		fmt.Printf("%s\n", item.ID())
	}
	for _, item := range f.ListNodes() {
		fmt.Printf("%s\n", item.ID())
	}
	for _, item := range f.ListEdges() {
		fmt.Printf("%d\n", item.ID())
	}
	//Output:
	//<nil> <nil> <nil>
	//<nil> <nil> n1 --> n2
	//
	//sg1
	//n1
	//n2
	//0
	//1
}

// Generating URLs for the mermaid live editor
func ExampleFlowchart_liveEditorLinks() {
	f := flowchart.NewFlowchart()
	f.AddEdge(f.AddNode("n1"), f.AddNode("n2"))
	// you can also use f.ViewInBrowser() to open the URL in browser directly
	fmt.Println(f.LiveURL())
	//Output: https://mermaidjs.github.io/mermaid-live-editor/#/view/eyJjb2RlIjoiZ3JhcGggVEJcbm4xW1wibjFcIl1cbm4yW1wibjJcIl1cbm4xIC0tXHUwMDNlIG4yXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9fQ==
}
