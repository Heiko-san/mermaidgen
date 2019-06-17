# mermaidgen

go package to generate mermaid code in an OOP way

# Quickstart

Here's a simple example that covers most of the functionality of this package
and therefore should be the quickest way to get started. 

``` go
package main

import (
    "fmt"
    m "github.com/Heiko-san/mermaidgen"
)

func main() {
    // create a new flowchart
    f := m.NewFlowchart()
    // set the direction (DirectionTopDown is the default)
    f.Direction = m.DirectionTopDown

    // define some styles if you want, there are styles for nodes and edges
    ns1 := f.NodeStyle("style1")
    ns1.Fill = m.ColorCyan
    ns1.Stroke = m.ColorBlack
    ns1.StrokeWidth = 1 //px
    ns1.StrokeDash = 5 //px
    //ns1.More for additional freestyle style definitions
    es1 := f.EdgeStyle("style1")
    es1.Stroke = m.ColorRed
    es1.StrokeWidth = 2 //px
    es1.StrokeDash = 5 //px
    //es1.More for additional freestyle style definitions

    // add some nodes and modify / style them
    n1 := f.AddNode("node1")
    // if no text is added the node's ID would be used as text
    n1.AddLines("this is node1", "textline 2")
    // default shape is NShapeRect
    n1.Shape = m.NShapeRoundRect
    // optional "click" line
    n1.Link = "http://www.example.com"
    // if this is not set, node.Link is used as tooltip
    n1.LinkText = "example"
    // add an optional styling
    n1.Style = ns1
    n2 := f.AddNode("node2")
    n2.AddLines("this is node2")
    n2.AddLines("textline 2")
    // you can also lookup an already defined style
    n2.Style = f.NodeStyle("style1")

    // add some edges and modify / style them
    e1 := f.AddEdge(n1, n2)
    // default shape is EShapeArrow
    e1.Shape = m.EShapeThickLine
    e1.AddLines("line1", "line2")
    e2 := f.AddEdge(n1, n2)
    e2.Style = es1
    e3 := f.AddEdge(n1, n1)
    e3.Style = es1

    // subgraphs are supported, too
    s1 := f.AddSubgraph("subgraph1")
    s1.Title = "A Subgraph"
    // and also nested subgraphs
    s2 := s1.AddSubgraph("subgraph2")
    s1.AddSubgraph("subgraph3")
    s2.AddNode("node3")
    // you can also lookup nodes and subgraphs again
    f.GetSubgraph("subgraph3").AddNode("node4")
    f.AddEdge(f.GetNode("node3"), f.GetNode("node4"))
    // Add functions will return nil if the ID already exists,
    // while Get funtions will return nil if ID doesn't exist,
    // AddEdge will always succeed since there are no IDs

    // you can also iterate over all items using f.ListSubgraphs, f.ListNodes
    // and f.ListEdges, which return slice-copies of the internal entity
    // storages containing pointers to the respective types

    // "render" the mermaid code by stringifying the flowchart 
    fmt.Print(f)
}
```
This generates the following mermaid code ...
``` plain
graph TB
classDef style1 stroke-width:1px,fill:#0ff,stroke:#000,stroke-dasharray:5px
node1("this is node1<br/>textline 2")
class node1 style1
click node1 "http://www.example.com" "example"
node2["this is node2<br/>textline 2"]
class node2 style1
subgraph A Subgraph
subgraph 
node3["node3"]
end
subgraph 
node4["node4"]
end
end
node1===|line1<br/>line2|node2
node1-->node2
linkStyle 1 stroke-width:2px,stroke:#f00,stroke-dasharray:5px
node1-->node1
linkStyle 2 stroke-width:2px,stroke:#f00,stroke-dasharray:5px
node3-->node4
```
... which renders to [this](https://mermaidjs.github.io/mermaid-live-editor/#/view/eyJjb2RlIjoiZ3JhcGggVEJcbmNsYXNzRGVmIHN0eWxlMSBzdHJva2Utd2lkdGg6MXB4LGZpbGw6IzBmZixzdHJva2U6IzAwMCxzdHJva2UtZGFzaGFycmF5OjVweFxubm9kZTEoXCJ0aGlzIGlzIG5vZGUxPGJyLz50ZXh0bGluZSAyXCIpXG5jbGFzcyBub2RlMSBzdHlsZTFcbmNsaWNrIG5vZGUxIFwiaHR0cDovL3d3dy5leGFtcGxlLmNvbVwiIFwiZXhhbXBsZVwiXG5ub2RlMltcInRoaXMgaXMgbm9kZTI8YnIvPnRleHRsaW5lIDJcIl1cbmNsYXNzIG5vZGUyIHN0eWxlMVxuc3ViZ3JhcGggQSBTdWJncmFwaFxuc3ViZ3JhcGggXG5ub2RlM1tcIm5vZGUzXCJdXG5lbmRcbnN1YmdyYXBoIFxubm9kZTRbXCJub2RlNFwiXVxuZW5kXG5lbmRcbm5vZGUxPT09fGxpbmUxPGJyLz5saW5lMnxub2RlMlxubm9kZTEtLT5ub2RlMlxubGlua1N0eWxlIDEgc3Ryb2tlLXdpZHRoOjJweCxzdHJva2U6I2YwMCxzdHJva2UtZGFzaGFycmF5OjVweFxubm9kZTEtLT5ub2RlMVxubGlua1N0eWxlIDIgc3Ryb2tlLXdpZHRoOjJweCxzdHJva2U6I2YwMCxzdHJva2UtZGFzaGFycmF5OjVweFxubm9kZTMtLT5ub2RlNCIsIm1lcm1haWQiOnsidGhlbWUiOiJkZWZhdWx0In19).
