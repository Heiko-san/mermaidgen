package mermaidgen

import (
	"fmt"
	"strings"
)

// An EdgeStyle is used to add CSS to an Edge. It renders to a linkStyle line
// for each Edge it is associated with. Note that linkStyles will override any
// effect from the Edge's shape defintion.
// Retrieve an instance of EdgeStyle via Flowchart's EdgeStyle method, do not
// create instances directly.
type EdgeStyle struct {
	id          string    // virtual ID for lookup
	Stroke      htmlColor // renders to something like stroke:#333
	StrokeWidth uint8     // renders to something like stroke-width:2px
	StrokeDash  uint8     // renders to something like stroke-dasharray:5px
	More        string    // more styles, e.g.: stroke:#333,stroke-width:1px
}

// ID provides access to the EdgeStyle's readonly field id.
func (es *EdgeStyle) ID() (id string) {
	return es.id
}

// String renders this graph element to a linkStyle line.
func (es *EdgeStyle) String() (renderedElement string) {
	styles := []string{
		fmt.Sprintf(`stroke-width:%dpx`, es.StrokeWidth),
	}
	if es.Stroke != "" {
		styles = append(styles, fmt.Sprintf(`stroke:%s`, es.Stroke))
	}
	if es.StrokeDash != 0 {
		styles = append(styles, fmt.Sprintf(`stroke-dasharray:%dpx`,
			es.StrokeDash))
	}
	if es.More != "" {
		styles = append(styles, es.More)
	}
	return fmt.Sprintf("linkStyle %s %s\n", "%d", strings.Join(styles, ","))
}
