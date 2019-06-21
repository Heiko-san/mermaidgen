package gantt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Heiko-san/mermaidgen/gantt"
)

// Accessing the readonly fields of a Task
func ExampleTask_privateFields() {
	g := gantt.NewGantt()
	s := g.AddSection("sect1")
	t, _ := s.AddTask("this_is_my_id")
	// get a copy of the id field
	id := t.ID()
	// access the top level Gantt
	gx := t.Gantt()
	// access this Task's Section
	sx := t.Section()
	fmt.Println(id, gx == g, sx == s)
	//Output:
	//this_is_my_id true true
}

// Different ways to set the fields of a Task
func ExampleTask_settingValues() {
	g := gantt.NewGantt()
	// Get a Task with only ID and default values, there is little chance for
	// an error here if you don't mess up the ID.
	t1, _ := g.AddTask("id1")
	// you can set the fields directly
	timestamp := time.Date(2019, 6, 20, 9, 15, 30, 0, time.UTC)
	t1.Start = &timestamp
	t1.Critical = true
	// to set Start (or After) and Duration there are helper functions that take
	// different values, see method description for details
	t2, _ := g.AddTask("id2")
	t3, _ := g.AddTask("id3")
	t4, _ := g.AddTask("id4")
	t5, _ := g.AddTask("id5")
	t6, _ := g.AddTask("id6")
	// start at time.Time
	t2.SetStart(&timestamp)
	t3.SetStart(timestamp)
	// start after Task
	t4.SetStart(t3)
	t5.SetStart("id4")
	// start at RFC3339 string
	t6.SetStart("2019-06-20T11:15:30+02:00")
	duration := time.Hour * 10
	end := timestamp.Add(time.Hour * 5)
	// duration by time.Time - Start (Start needs to be defined, not After)
	t2.SetDuration(&end)
	t2.SetDuration(end)
	// duration by time.Duration
	t3.SetDuration(&duration)
	t3.SetDuration(duration)
	// copy duration from other Task
	t4.SetDuration(t1)
	t5.SetDuration("id3")
	// duration string
	t6.SetDuration("12h30m30s")
	// you can provide values directly to AddTask method
	// (id, Title, SetDuration(), SetStart(), Critical, Active, Done)
	t7, _ := g.AddTask("id7", "A Task", time.Hour*20, timestamp, true, true, true)
	// finally you can copy the settings from other Tasks
	t8, _ := g.AddTask("id8")
	t8.CopyFields(t7)
	fmt.Print(g)
	//Output:
	//gantt
	//dateFormat YYYY-MM-DDTHH:mm:ssZ
	//id1 : crit, id1, 2019-06-20T09:15:30Z, 1d
	//id2 : id2, 2019-06-20T09:15:30Z, 18000s
	//id3 : id3, 2019-06-20T09:15:30Z, 36000s
	//id4 : id4, after id3, 1d
	//id5 : id5, after id4, 36000s
	//id6 : id6, 2019-06-20T11:15:30+02:00, 45030s
	//A Task : crit, active, done, id7, 2019-06-20T09:15:30Z, 72000s
	//A Task : crit, active, done, id8, 2019-06-20T09:15:30Z, 72000s
}

// The creation of Tasks may yield errors
func ExampleTask_errorHandling() {
	g := gantt.NewGantt()
	// if there are parameter errors, no Task is created, an error is returned
	t1, err := g.AddTask("id1", "my title", "1h50xyz")
	fmt.Println(t1, err)
	t1, err = g.AddTask("id1", "my title", "1h50m", "foobar")
	fmt.Println(t1, err)
	// same applies if the requested ID already exists
	t1, err = g.AddTask("id1")
	t2, err := g.AddTask("id1")
	fmt.Println(t2, err)
	// or if an invalid ID is provided
	t2, err = g.AddTask("foo bar baz")
	fmt.Println(t2, err)
	//Output:
	//<nil> SetDuration: "1h50xyz" is neither a valid duration nor Task ID
	//<nil> SetStart: "foobar" is neither RFC3339 nor a valid Task ID
	//<nil> id already exists
	//<nil> invalid id
}

func TestErrors(t *testing.T) {
	g := gantt.NewGantt()
	if _, err := g.AddTask("id1", 2); err == nil {
		t.Errorf("no error returned: Title invalid type")
	}
	if _, err := g.AddTask("id1", "", false); err == nil {
		t.Errorf("no error returned: Duration invalid type")
	}
	if _, err := g.AddTask("id1", "", time.Now()); err == nil {
		t.Errorf("no error returned: Start=nil")
	}
	if _, err := g.AddTask("id1", "", "1h", false); err == nil {
		t.Errorf("no error returned: Start invalid type")
	}
	if _, err := g.AddTask("id1", "", "1h", time.Now(), 5); err == nil {
		t.Errorf("no error returned: Invalid flags")
	}
	if _, err := g.AddTask("id1", "", "1h", time.Now(), true, 5); err == nil {
		t.Errorf("no error returned: Invalid flags")
	}
	if _, err := g.AddTask("id1", "", "1h", time.Now(), true, true, 5); err == nil {
		t.Errorf("no error returned: Invalid flags")
	}
}

func TestCopy(t *testing.T) {
	g := gantt.NewGantt()
	s := g.AddSection("s")
	t0, _ := g.AddTask("id0")
	t1, _ := s.AddTask("id1", "title", "1h")
	t1.Active = true
	t1.After = t0
	t2, _ := g.AddTask("id2")
	t2.CopyFields(t1)
	if t1.Critical != t2.Critical {
		t.Fail()
	}
	if t1.Active != t2.Active {
		t.Fail()
	}
	if t1.Done != t2.Done {
		t.Fail()
	}
	if t2.After != t0 {
		t.Fail()
	}
	if t1.Duration == t2.Duration {
		t.Fail()
	}
	if *t1.Duration != *t2.Duration {
		t.Fail()
	}
	if t2.Start != nil {
		t.Fail()
	}
	if t1.Title != t2.Title {
		t.Fail()
	}
	if t1.ID() == t2.ID() {
		t.Fail()
	}
	if t1.Section() == t2.Section() {
		t.Fail()
	}
	t1.SetStart(time.Now())
	t2.CopyFields(t1)
	if t1.Start == t2.Start {
		t.Fail()
	}
	if *t1.Start != *t2.Start {
		t.Fail()
	}
}
