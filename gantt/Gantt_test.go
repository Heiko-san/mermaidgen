package gantt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Heiko-san/mermaidgen/gantt"
)

// Defining and rendering a gantt diagram
func ExampleGantt_basics() {
	// create a gantt diagram
	g, _ := gantt.NewGantt()
	// set properties for the different elements
	g.Title = "My diagram"
	// add some tasks to it directly
	t1, _ := g.AddTask("t1")
	t1.SetStart(time.Date(2019, 6, 20, 9, 15, 30, 0, time.UTC))
	t1.Critical = true
	// or add sections with tasks
	s1, _ := g.AddSection("s1")
	s1.AddTask("t2", "section1 task", "10h30m", t1)
	s2, _ := g.AddSection("s2")
	s2.AddTask("t3", "section2 task", "20h", t1)
	// stringify renders the mermaid code
	fmt.Print(g)
	// you can also look up already defined Sections and Tasks
	s1x := g.GetSection("s1")
	t1x := g.GetTask("t1")
	fmt.Println("they are the same:", s1 == s1x, t1 == t1x)
	//Output:
	//gantt
	//dateFormat YYYY-MM-DDTHH:mm:ssZ
	//title My diagram
	//t1 : crit, t1, 2019-06-20T09:15:30Z, 1d
	//section s1
	//section1 task : t2, after t1, 37800s
	//section s2
	//section2 task : t3, after t1, 72000s
	//they are the same: true true
}

// Properties of a gantt diagram
func ExampleGantt_settings() {
	// setting Title and AxisFormat on object creation
	g, _ := gantt.NewGantt("This is my title", "%H:%M")
	// setting the Title afterwards
	g.Title = "This is my title"
	// setting AxisFormat afterwards using constants
	g.AxisFormat = gantt.FormatTime24
	fmt.Print(g)
	//Output:
	//gantt
	//dateFormat YYYY-MM-DDTHH:mm:ssZ
	//axisFormat %H:%M
	//title This is my title
}

// Iterate over the Tasks and Sections of a Gantt diagram
func ExampleGantt_iterateSectionsAndTasks() {
	g, _ := gantt.NewGantt()
	// add some local tasks to Gantt
	g.AddTask("t1")
	g.AddTask("t2")
	// add sections
	s1, _ := g.AddSection("This is my section")
	s2, _ := g.AddSection("Another section")
	// add some tasks to them
	s1.AddTask("a3")
	s2.AddTask("b4")
	// iterate over Sections
	fmt.Println("all sections")
	for _, s := range g.ListSections() {
		fmt.Println(s.ID())
	}
	// iterate over Tasks
	fmt.Println("this Gantt's local tasks")
	for _, t := range g.ListLocalTasks() {
		fmt.Println(t.ID())
	}
	fmt.Println("all tasks")
	for _, t := range g.ListTasks() {
		fmt.Println(t.ID())
	}
	//Output:
	//all sections
	//This is my section
	//Another section
	//this Gantt's local tasks
	//t1
	//t2
	//all tasks
	//a3
	//b4
	//t1
	//t2
}

// Generating URLs for the mermaid live editor
func ExampleGantt_liveEditorLinks() {
	g, _ := gantt.NewGantt()
	timestamp := time.Date(2019, 6, 20, 9, 15, 30, 0, time.UTC)
	g.AddTask("t1", "a task", "1h", timestamp)
	g.AddTask("t2", "another task", "2h")
	// you can also use g.ViewInBrowser() to open the URL in browser directly
	fmt.Println(g.LiveURL())
	//Output: https://mermaidjs.github.io/mermaid-live-editor/#/view/eyJjb2RlIjoiZ2FudHRcbmRhdGVGb3JtYXQgWVlZWS1NTS1ERFRISDptbTpzc1pcbmEgdGFzayA6IHQxLCAyMDE5LTA2LTIwVDA5OjE1OjMwWiwgMzYwMHNcbmFub3RoZXIgdGFzayA6IDcyMDBzXG4iLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9fQ==
}

// The creation of Gantts, Sections and Tasks may yield errors
func ExampleGantt_errorHandling() {
	// parameter errors will be returned and no Gantt will be created
	g1, err := gantt.NewGantt(5, gantt.FormatDate)
	fmt.Println(g1, err)
	g2, err := gantt.NewGantt("title", 5)
	fmt.Println(g2, err)
	g, _ := gantt.NewGantt()
	// same applies to Task creation
	t1, err := g.AddTask("id1", "my title", "1h50xyz")
	fmt.Println(t1, err)
	t1, err = g.AddTask("id1", "my title", "1h50m", "foobar")
	fmt.Println(t1, err)
	// Sections have no initializers that may fail, however errors will be
	// returned if you try to create a duplicate ID
	s1, _ := g.AddSection("s1")
	s2, err := g.AddSection("s1")
	fmt.Println(s2, err)
	// same applies to Task creation
	t1, err = s1.AddTask("id1")
	t2, err := g.AddTask("id1")
	fmt.Println(t2, err)
	// or if an invalid ID is provided for Task creation
	t2, err = g.AddTask("foo bar baz")
	fmt.Println(t2, err)
	//Output:
	//<nil> value for Title was no string
	//<nil> value for AxisFormat was no axisFormat
	//<nil> SetDuration: "1h50xyz" is neither a valid duration nor Task ID
	//<nil> SetStart: "foobar" is neither RFC3339 nor a valid Task ID
	//<nil> id already exists
	//<nil> id already exists
	//<nil> invalid id
}

func TestGanttAddDuplicateTask(t *testing.T) {
	g, _ := gantt.NewGantt()
	t1, _ := g.AddTask("t1")
	t2, err := g.AddTask("t1")
	assert(t, t1 != t2)
	assert(t, t2 == nil)
	assert(t, err != nil)
}

func TestGanttAddDuplicateSection(t *testing.T) {
	g, _ := gantt.NewGantt()
	s1, _ := g.AddSection("s1")
	s2, err := g.AddSection("s1")
	assert(t, s1 != s2)
	assert(t, s2 == nil)
	assert(t, err != nil)
}
