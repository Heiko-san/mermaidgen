package flowchart_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/flowchart"
)

// Working with Edges
func ExampleEdge() {
	f := flowchart.NewFlowchart()
	n1 := f.AddNode("n1")
	n2 := f.AddNode("n2")
	// define some Edges
	e1 := f.AddEdge(n1, n2)
	e2 := f.AddEdge(n1, n2)
	// you can access and modify the Node members afterwards
	e2.From = n2
	e2.To = n1
	// define the Edge shape (EdgeStyles may override the shape)
	e1.Shape = flowchart.EShapeDottedLine
	// add text (you can also directly access e1.Text slice)
	e1.AddLines("first line", "second line")
	e1.AddLines("third line")
	// CSS styling (see EdgeStyle for more details)
	e2.Style = f.EdgeStyle("es1")
	// Previously defined Edges can be looked up
	f.GetEdge(1).Style.Stroke = flowchart.ColorCyan
	fmt.Print(f)
	//Output:
	//graph TB
	//n1["n1"]
	//n2["n2"]
	//n1 -.-|"first line<br/>second line<br/>third line"| n2
	//n2 --> n1
	//linkStyle 1  stroke:#0ff
}

// Accessing the readonly fields of an Edge
func ExampleEdge_privateFields() {
	f := flowchart.NewFlowchart()
	n1 := f.AddNode("n1")
	n2 := f.AddNode("n2")
	e1 := f.AddEdge(n1, n2)
	e2 := f.AddEdge(n1, n2)
	// the ID of an Edge is actually the order of the Edge creation
	fmt.Println(e2.ID(), e1.ID())
	//Output: 1 0
}
