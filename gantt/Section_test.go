package gantt_test

import (
	"fmt"
	"testing"

	"github.com/Heiko-san/mermaidgen/gantt"
)

// Iterate over the Tasks of a Section
func ExampleSection_iterateSectionTasks() {
	g, _ := gantt.NewGantt()
	g.AddTask("t1")
	// add sections
	s1, _ := g.AddSection("This is my section")
	s2, _ := g.AddSection("Another section")
	// add some tasks to them
	s1.AddTask("t2")
	s1.AddTask("t3")
	s2.AddTask("t4")
	// iterate over s1 Tasks
	for _, t := range s1.ListLocalTasks() {
		fmt.Println(t.ID())
	}
	//Output:
	//t2
	//t3
}

// Accessing the readonly fields of a Section
func ExampleSection_privateFields() {
	g, _ := gantt.NewGantt()
	s, _ := g.AddSection("this_is_my_id")
	// get a copy of the id field
	id := s.ID()
	// access the top level Gantt
	gx := s.Gantt()
	fmt.Println(id, gx == g)
	//Output:
	//this_is_my_id true
}

func TestSection_addDuplicateTask(t *testing.T) {
	g, _ := gantt.NewGantt()
	s, _ := g.AddSection("s")
	t1, _ := s.AddTask("t1")
	t2, err := s.AddTask("t1")
	assert(t, t1 != t2)
	assert(t, t2 == nil)
	assert(t, err != nil)
}
