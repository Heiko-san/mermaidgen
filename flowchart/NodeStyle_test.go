package flowchart_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/flowchart"
)

// Working with NodeStyles
func ExampleNodeStyle() {
	f := flowchart.NewFlowchart()
	n1 := f.AddNode("n1")
	n2 := f.AddNode("n2")
	// create one or more styles
	ns1 := f.NodeStyle("ns1")
	ns2 := f.NodeStyle("ns2")
	// and assign them to Nodes
	n1.Style = ns1
	// you can also lookup previously defined NodeStyles
	n2.Style = f.NodeStyle("ns2")
	// there are some CSS shortcuts
	ns2.Fill = flowchart.ColorYellow
	ns2.Stroke = flowchart.ColorBlue
	ns2.StrokeWidth = 2
	ns2.StrokeDash = 5
	// but you can also add additional styles, just in case
	ns2.More = "font-size:20px"
	fmt.Print(f)
	//Output:
	//graph TB
	//classDef ns1 stroke-width:1px
	//classDef ns2 fill:#ff0,stroke:#00f,stroke-width:2px,stroke-dasharray:5px,font-size:20px
	//n1["n1"]
	//class n1 ns1
	//n2["n2"]
	//class n2 ns2
}

// Accessing the readonly fields of a NodeStyle
func ExampleNodeStyle_privateFields() {
	f := flowchart.NewFlowchart()
	ns := f.NodeStyle("this_is_my_id")
	// get a copy of the id field
	fmt.Println(ns.ID())
	//Output: this_is_my_id
}
