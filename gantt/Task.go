package gantt

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Task represents gantt tasks that can be added to Sections or the Gantt
// diagram itself. Create an instance of Task via Gantt's or Section's AddTask
// method, do not create instances directly. Already defined IDs can be looked
// up via Gantt's GetTask method or iterated over via its ListTasks method.
type Task struct {
	id       string         // Task ID
	gantt    *Gantt         // The top level Gantt diagram
	section  *Section       // The Section this Task belongs to
	Title    string         // Title of the Task, if not set, ID is used
	Start    *time.Time     // Time when the Task starts (Start wins over After)
	After    *Task          // Task after which this Task starts
	Duration *time.Duration // Duration of the Task (the absolute value is used)
	Critical bool           // The crit flag
	Active   bool           // The active flag
	Done     bool           // The done flag
}

// ID provides access to the Task's readonly field id.
func (t *Task) ID() (id string) {
	return t.id
}

// Gantt provides access to the top level Gantt diagram to be able to access
// Adder, Getter and Lister methods.
func (t *Task) Gantt() (topLevel *Gantt) {
	return t.gantt
}

// Section provides access to the Section that contains this Task to be able to
// access Adder methods. If this Task is contained in the top level Gantt
// itself, nil is returned.
func (t *Task) Section() (containingSection *Section) {
	return t.section
}

// String renders this diagram element to a task definition line.
func (t *Task) String() (renderedElement string) {
	title := t.Title
	if title == "" {
		title = t.id
	}
	tokens := []string{}
	// flags
	if t.Critical {
		tokens = append(tokens, "crit")
	}
	if t.Active {
		// active > done when rendering
		tokens = append(tokens, "active")
	}
	if t.Done {
		tokens = append(tokens, "done")
	}
	// functional
	if t.Start != nil {
		// id without start statement breaks syntax
		tokens = append(tokens, t.id)
		tokens = append(tokens, t.Start.Format(time.RFC3339))
	} else if t.After != nil {
		tokens = append(tokens, t.id)
		tokens = append(tokens, "after "+t.After.id)
	}
	duration := "1d"
	if t.Duration != nil {
		duration = fmt.Sprintf("%ds", int(math.Abs(t.Duration.Seconds())))
	}
	tokens = append(tokens, duration)
	text := fmt.Sprintf("%s : %s\n", title, strings.Join(tokens, ", "))
	return text
}

// SetStart takes a time.Time or a pointer to it, a Task pointer or a string
// that represents an existing Task ID or a RFC3339 time definition and sets
// this Task's Start or After field from that information. An error is returned
// if the given type is not supported or the string can't be parsed.
func (t *Task) SetStart(start interface{}) (err error) {
	switch tStart := start.(type) {
	case *time.Time:
		t.Start = tStart
	case *Task:
		t.After = tStart
		// time > after -> unset time
		t.Start = nil
	case time.Time:
		t.Start = &tStart
	case string:
		if task := t.gantt.GetTask(tStart); task != nil {
			t.After = task
			t.Start = nil
		} else {
			x, err := time.Parse(time.RFC3339, tStart)
			if err != nil {
				return fmt.Errorf(
					`string "%s" is neither RFC3339 nor a valid Task ID`,
					tStart)
			}
			t.Start = &x
		}
	default:
		return fmt.Errorf("unsupported type: %T", start)
	}
	return nil
}

// Helperfunction to deduplicate code.
func (t *Task) setDurationFromTime(endTime *time.Time) (err error) {
	if t.Start != nil {
		duration := endTime.Sub(*t.Start)
		t.Duration = &duration
		return nil
	}
	return fmt.Errorf("can't calculate duration, Start is not defined")
}

// Helperfunction to deduplicate code.
func (t *Task) setDurationFromTask(task *Task) {
	if task.Duration == nil {
		t.Duration = nil
	} else {
		*t.Duration = *task.Duration
	}
}

// SetDuration takes a time.Duration or a pointer to it, a time.Time or a
// pointer to it, a Task pointer or a string that represents an existing Task ID
// or a parsable duration definition and sets this Task's Duration field from
// that information. If this information represents a time.Time, the difference
// to Start is calculated, if it represents a Task, that Task's Duration is
// copied. An error is returned if the given type is not supported, Start is
// undefined for a Time definition or the string can't be parsed.
func (t *Task) SetDuration(duration interface{}) (err error) {
	switch tDuration := duration.(type) {
	case *time.Duration:
		t.Duration = tDuration
	case time.Duration:
		t.Duration = &tDuration
	case *time.Time:
		return t.setDurationFromTime(tDuration)
	case time.Time:
		return t.setDurationFromTime(&tDuration)
	case *Task:
		t.setDurationFromTask(tDuration)
	case string:
		if task := t.gantt.GetTask(tDuration); task != nil {
			t.setDurationFromTask(task)
		} else {
			x, err := time.ParseDuration(tDuration)
			if err != nil {
				return fmt.Errorf(
					`string "%s" is neither a valid duration nor Task ID`,
					tDuration)
			}
			t.Duration = &x
		}
	default:
		return fmt.Errorf("unsupported type: %T", duration)
	}
	return nil
}
