package mermaidgen

import (
	"fmt"
)

////////// Subgraph ////////////////////////////////////////////////////////////

type subgraph struct {
	// ro fields
	id						string			// virtual ID for lookup
	flowchart				*flowchart
	items					[]graphItem
	// rw fields
	Title					string
}

func (self *subgraph) ID() string {
	return self.id
}

func (self *subgraph) Flowchart() *flowchart {
	return self.flowchart
}

// render
func (self *subgraph) renderGraph() string {
 	text := fmt.Sprintf("subgraph %s\n", self.Title)
	for _, item := range self.items {
		text += item.renderGraph()
	}
	text += "end\n"
	return text
}

// stringify
func (self *subgraph) String() string {
	return self.renderGraph()
}

////////// add Items ///////////////////////////////////////////////////////////

func (self *subgraph) AddSubgraph(id string) *subgraph {
	_, found := self.flowchart.subgraphs[id]
	if found {
		// if already exists -> nil
		return nil
	} else {
		s := &subgraph{id: id, flowchart: self.flowchart}
		self.flowchart.subgraphs[id] = s
		self.items = append(self.items, s)
		return s
	}
}

func (self *subgraph) AddNode(id string) *node {
	_, found := self.flowchart.nodes[id]
	if found {
		// if already exists -> nil
		return nil
	} else {
		n := &node{id: id, Shape: NShapeRect}
		self.flowchart.nodes[id] = n
		self.items = append(self.items, n)
		return n
	}
}