package gantt_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/gantt"
)

// Working with Sections
func ExampleSection() {
	g := gantt.NewGantt()
	fmt.Print(g)
	//Output:
	//gantt
	//dateFormat YYYY-MM-DDTHH:mm:ssZ
}
