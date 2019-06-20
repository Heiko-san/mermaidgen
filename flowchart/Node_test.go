package flowchart_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/flowchart"
)

// Working with Nodes
func ExampleNode() {
	f := flowchart.NewFlowchart()
	// define some Nodes
	n1 := f.AddNode("n1")
	n2 := f.AddNode("n2")
	// define the Node shape
	n1.Shape = flowchart.NShapeRoundRect
	// add text (you can also directly access n1.Text slice)
	n1.AddLines("first line", "second line")
	n1.AddLines("third line")
	// you may add a link
	n2.Link = "http://www.example.com"
	n2.LinkText = "tooltip"
	// CSS styling (see NodeStyle for more details)
	n2.Style = f.NodeStyle("ns1")
	// Previously defined Nodes can be looked up
	f.GetNode("n2").Style.Fill = flowchart.ColorCyan
	fmt.Print(f)
	//Output:
	//graph TB
	//classDef ns1 fill:#0ff
	//n1("first line<br/>second line<br/>third line")
	//n2["n2"]
	//class n2 ns1
	//click n2 "http://www.example.com" "tooltip"
}

// Accessing the readonly fields of a Node
func ExampleNode_privateFields() {
	f := flowchart.NewFlowchart()
	n1 := f.AddNode("this_is_my_id")
	// get a copy of the id field
	fmt.Println(n1.ID())
	//Output: this_is_my_id
}
