package gantt_test

import (
	"fmt"

	"github.com/Heiko-san/mermaidgen/gantt"
)

// Working with Sections
func ExampleTask() {
	g := gantt.NewGantt()
	fmt.Print(g)
	//Output:
	//gantt
	//dateFormat YYYY-MM-DDTHH:mm:ssZ
}
