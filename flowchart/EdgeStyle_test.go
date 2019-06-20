package flowchart_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/flowchart"
)

// Working with EdgeStyles
func ExampleEdgeStyle() {
	f := flowchart.NewFlowchart()
	n1 := f.AddNode("n1")
	n2 := f.AddNode("n2")
	e1 := f.AddEdge(n1, n2)
	e2 := f.AddEdge(n1, n2)
	// create one or more styles
	es1 := f.EdgeStyle("es1")
	es2 := f.EdgeStyle("es2")
	// and assign them to Edges
	e1.Style = es1
	// you can also lookup previously defined EdgeStyles
	e2.Style = f.EdgeStyle("es2")
	// there are some CSS shortcuts
	es2.Stroke = flowchart.ColorRed
	es2.StrokeWidth = 2
	es2.StrokeDash = 5
	// but you can also add additional styles, just in case
	es2.More = "font-size:20px"
	fmt.Print(f)
	//Output:
	//graph TB
	//n1["n1"]
	//n2["n2"]
	//n1 --> n2
	//linkStyle 0  stroke-width:1px
	//n1 --> n2
	//linkStyle 1  stroke:#f00,stroke-width:2px,stroke-dasharray:5px,font-size:20px
}

// Accessing the readonly fields of an EdgeStyle
func ExampleEdgeStyle_privateFields() {
	f := flowchart.NewFlowchart()
	es := f.EdgeStyle("this_is_my_id")
	// get a copy of the id field
	fmt.Println(es.ID())
	//Output: this_is_my_id
}

// Defining default linkStyles and curve interpolation
func ExampleEdgeStyle_defaultStyle() {
	f := flowchart.NewFlowchart()
	f.DefaultEdgeStyle = f.EdgeStyle("my_style_1")
	f.DefaultEdgeStyle.Interpolation = flowchart.InterpolationBasis
	fmt.Print(f)
	//Output:
	//graph TB
	//linkStyle default interpolate basis
}
