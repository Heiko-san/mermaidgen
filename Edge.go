package mermaidgen

import (
	"fmt"
	"strings"
)

////////// EdgeShape ///////////////////////////////////////////////////////////

type edgeShape string
const (
	EShapeArrow 		edgeShape =  `-->`
	EShapeDottedArrow	edgeShape =  `-.->`
	EShapeThickArrow	edgeShape =  `==>`
	EShapeLine 			edgeShape =  `---`
	EShapeDottedLine	edgeShape =  `-.-`
	EShapeThickLine		edgeShape =  `===`
)

////////// Edge ////////////////////////////////////////////////////////////////

type edge struct {
	// ro fields
	id				int
	// rw fields
	From 			*node
	To 				*node
	Shape 			edgeShape
	Text 			[]string
	Style 			*edgeStyle
}

func (self *edge) ID() int {
	return self.id
}

// render / stringify
func (self *edge) String() string {
	line := string(self.Shape)
	if len(self.Text) > 0 {
		line += fmt.Sprintf(`|%s|`, strings.Join(self.Text, "<br/>"))
	}
	text := self.From.id + line + self.To.id + "\n"
	if self.Style != nil {
		text += fmt.Sprintf(self.Style.String(), self.id)
	}
	return text
}

func (self *edge) AddLines(lines ...string) {
	self.Text = append(self.Text, lines...)
}