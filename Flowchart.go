package mermaidgen
// https://mermaidjs.github.io/flowchart.html

import (
	"fmt"
)

////////// ChartDirection //////////////////////////////////////////////////////

type chartDirection string
const (
	DirectionTopDown 		chartDirection = `TB`
	DirectionBottomUp 		chartDirection = `BT`
	DirectionRightLeft 		chartDirection = `RL`
	DirectionLeftRight 		chartDirection = `LR`
)

////////// GraphItem ///////////////////////////////////////////////////////////

// interface to define what can be an "item" to a Flowchart/Subgraph
type graphItem interface{
	renderGraph() 			string
}

////////// Flowchart ///////////////////////////////////////////////////////////

type flowchart struct {
	// ro fields
	nodeStyles 				map[string]*nodeStyle
	edgeStyles 				map[string]*edgeStyle
	subgraphs 				map[string]*subgraph
	nodes 					map[string]*node
	edges 					[]*edge
	items					[]graphItem
	// rw fields
	Direction				chartDirection
}

func NewFlowchart() *flowchart {
	f := &flowchart{}
	f.Direction = DirectionTopDown
	f.nodeStyles = make(map[string]*nodeStyle)
	f.edgeStyles = make(map[string]*edgeStyle)
	f.subgraphs = make(map[string]*subgraph)
	f.nodes = make(map[string]*node)
	return f
}

// render / stringify
func (self *flowchart) String() string {
	text := fmt.Sprintf("graph %s\n", self.Direction)
	for _, s := range self.nodeStyles {
		text += s.String()
	}
	for _, item := range self.items {
		text += item.renderGraph()
	}
	for _, e := range self.edges {
		text += e.String()
	}
	return text
}

////////// add & get Styles ////////////////////////////////////////////////////

func (self *flowchart) NodeStyle(id string) *nodeStyle {
	s, found := self.nodeStyles[id]
	if !found {
		s = &nodeStyle{id: id, StrokeWidth: 1}
		self.nodeStyles[id] = s
	}
	return s
}

func (self *flowchart) EdgeStyle(id string) *edgeStyle {
	s, found := self.edgeStyles[id]
	if !found {
		s = &edgeStyle{id: id, StrokeWidth: 1}
		self.edgeStyles[id] = s
	}
	return s
}

////////// add Items ///////////////////////////////////////////////////////////

func (self *flowchart) AddSubgraph(id string) *subgraph {
	_, found := self.subgraphs[id]
	if found {
		// if already exists -> nil
		return nil
	} else {
		s := &subgraph{id: id, flowchart: self}
		self.subgraphs[id] = s
		self.items = append(self.items, s)
		return s
	}
}

func (self *flowchart) AddNode(id string) *node {
	_, found := self.nodes[id]
	if found {
		// if already exists -> nil
		return nil
	} else {
		n := &node{id: id, Shape: NShapeRect}
		self.nodes[id] = n
		self.items = append(self.items, n)
		return n
	}
}

func (self *flowchart) AddEdge(from *node, to *node) *edge {
	e := &edge{From: from, To: to, Shape: EShapeArrow}
	self.edges = append(self.edges, e)
	e.id = len(self.edges) - 1
	return e
}

////////// get Items ///////////////////////////////////////////////////////////

func (self *flowchart) GetSubgraph(id string) *subgraph {
	// if not found -> nil
	return self.subgraphs[id]
}

func (self *flowchart) GetNode(id string) *node {
	// if not found -> nil
	return self.nodes[id]
}

func (self *flowchart) GetEdge(index int) *edge {
	if len(self.edges) <= index || index < 0 {
		return nil
	}
	return self.edges[index]
}

////////// list Items //////////////////////////////////////////////////////////

func (self *flowchart) ListSubgraphs() []*subgraph {
    values := make([]*subgraph, 0, len(self.subgraphs))
    for _, v := range self.subgraphs {
        values = append(values, v)
    }
	return values
}

func (self *flowchart) ListNodes() []*node {
	values := make([]*node, 0, len(self.nodes))
    for _, v := range self.nodes {
        values = append(values, v)
    }
	return values
}

func (self *flowchart) ListEdges() []*edge {
	e := make([]*edge, len(self.edges))
	copy(e, self.edges)
	return e
}