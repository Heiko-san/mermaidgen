package mermaidgen

import (
	"fmt"
	"strings"
)

////////// NodeStyle ///////////////////////////////////////////////////////////

type edgeStyle struct {
	// ro fields
	id					string		// virtual ID for lookup
	// rw fields
	Stroke 				htmlColor	// stroke:#333
	StrokeWidth			uint8		// stroke-width:1px
	StrokeDash			uint8		// stroke-dasharray:0px
	More				string		// freestyle text definitions
}

func (self *edgeStyle) ID() string {
	return self.id
}

// render / stringify
func (self *edgeStyle) String() string {
	styles := []string{
		fmt.Sprintf(`stroke-width:%dpx`, self.StrokeWidth),
	}
	if self.Stroke != "" {
		styles = append(styles, fmt.Sprintf(`stroke:%s`, self.Stroke))
	}
	if self.StrokeDash != 0 {
		styles = append(styles, fmt.Sprintf(`stroke-dasharray:%dpx`,
			self.StrokeDash))
	}
	if self.More != "" {
		styles = append(styles, self.More)
	}
	return fmt.Sprintf("linkStyle %s %s\n", "%d", strings.Join(styles,","))
}