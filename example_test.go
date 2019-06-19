package mermaidgen_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen"
)

// Working with Subgraphs
func ExampleSubgraph() {
	f := mermaidgen.NewFlowchart()
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
