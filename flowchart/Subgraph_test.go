package flowchart_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/flowchart"
)

// Working with Subgraphs
func ExampleSubgraph() {
	f := flowchart.NewFlowchart()
	// add some Subgraphs
	sg1 := f.AddSubgraph("sg1")
	sg1.Title = "vpc-123"
	sg2 := sg1.AddSubgraph("sg2")
	sg2.Title = "AZ a"
	// oops we forgot to get the reference ...
	sg1.AddSubgraph("sg3")
	// ... but we can also look it up
	sg3 := f.GetSubgraph("sg3")
	sg3.Title = "AZ b"
	// add some Nodes to different Subgraphs
	f.AddEdge(sg2.AddNode("i-123"), sg2.AddNode("mydb"))
	f.AddEdge(sg3.AddNode("i-456"), f.GetNode("mydb"))
	fmt.Print(f)
	//Output:
	//graph TB
	//subgraph vpc-123
	//subgraph AZ a
	//i-123["i-123"]
	//mydb["mydb"]
	//end
	//subgraph AZ b
	//i-456["i-456"]
	//end
	//end
	//i-123 --> mydb
	//i-456 --> mydb
}

// Accessing the readonly fields of a Subgraph
func ExampleSubgraph_privateFields() {
	f := flowchart.NewFlowchart()
	sg := f.AddSubgraph("this_is_my_id")
	sg.Title = "this is my title"
	// get a copy of the id field
	id := sg.ID()
	// get a pointer to the top level Flowchart
	fx := sg.Flowchart()
	fmt.Println(id, f == fx)
	//Output: this_is_my_id true
}

// Adding the same ID multiple times won't work
func ExampleSubgraph_addDuplicate() {
	f := flowchart.NewFlowchart()
	f.AddNode("myid1")
	// myid1 can be used since we deal with a different type
	sg := f.AddSubgraph("myid1")
	// myid1 can't be used since for both types it was already defined
	n := sg.AddNode("myid1")
	s := sg.AddSubgraph("myid1")
	fmt.Println(n, s)
	//Output: <nil> <nil>
}
