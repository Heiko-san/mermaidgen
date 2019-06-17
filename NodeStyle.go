package mermaidgen

import (
	"fmt"
	"strings"
)

////////// NodeStyle ///////////////////////////////////////////////////////////

type nodeStyle struct {
	// ro fields
	id					string
	// rw fields
	Fill 				htmlColor	// fill:#f9f
	Stroke 				htmlColor	// stroke:#333
	StrokeWidth			uint8		// stroke-width:1px
	StrokeDash			uint8		// stroke-dasharray:0px
	More				string		// freestyle text definitions
}

func (self *nodeStyle) ID() string {
	return self.id
}

// render / stringify
func (self *nodeStyle) String() string {
	styles := []string{
		fmt.Sprintf(`stroke-width:%dpx`, self.StrokeWidth),
	}
	if self.Fill != "" {
		styles = append(styles, fmt.Sprintf(`fill:%s`, self.Fill))
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
	return fmt.Sprintf("classDef %s %s\n", self.id, strings.Join(styles,","))
}