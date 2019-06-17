package mermaidgen

import (
	"fmt"
	"strings"
)

////////// NodeShape ///////////////////////////////////////////////////////////

type nodeShape string
const (
	NShapeRect 			nodeShape =  `["%s"]`
	NShapeRoundRect 	nodeShape =  `("%s")`
	NShapeCircle 		nodeShape = `(("%s"))`
	NShapeRhombus 		nodeShape =  `{"%s"}`
	NShapeFlag 			nodeShape =  `>"%s"]`
)

////////// Node ////////////////////////////////////////////////////////////////

type node struct {
	// ro fields
	id					string
	// rw fields
	Shape				nodeShape
	Text 				[]string
	Link				string
	LinkText			string
	Style 				*nodeStyle
}

func (self *node) ID() string {
	return self.id
}

// render
func (self *node) renderGraph() string {
	textbox := self.id
	if len(self.Text) > 0 {
		textbox = strings.Join(self.Text,"<br/>")
	}
	text := self.id + fmt.Sprintf(string(self.Shape), textbox) + "\n"
	if self.Style != nil {
		text += fmt.Sprintf("class %s %s\n", self.id, self.Style.id)
	}
	if self.Link != "" {
		linktxt := self.Link
		if self.LinkText != "" {
			linktxt = self.LinkText
		}
		text += fmt.Sprintf("click %s \"%s\" \"%s\"\n",
			self.id, self.Link, linktxt)
	}
	return text
}

// stringify
func (self *node) String() string {
	return self.renderGraph()
}

// body text lines
func (self *node) AddLines(lines ...string) {
	self.Text = append(self.Text, lines...)
}