package gantt

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

// IsValidID is used to check if Task IDs are valid: IsValidID(string) bool.
// These limitations don't apply to Section IDs which are Section's titles at
// the same time.
var IsValidID = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString

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

// Private constructor for use in Add-functions.
func taskNew(i string, g *Gantt, s *Section, p []interface{}) (*Task, error) {
	if !IsValidID(i) {
		return nil, fmt.Errorf("invalid id")
	}
	t := &Task{id: i, gantt: g, section: s}
	switch l, ok := len(p), false; {
	case l > 5:
		t.Done, ok = p[5].(bool)
		if !ok {
			return nil, fmt.Errorf("value for Done was no bool")
		}
		fallthrough
	case l > 4:
		t.Active, ok = p[4].(bool)
		if !ok {
			return nil, fmt.Errorf("value for Active was no bool")
		}
		fallthrough
	case l > 3:
		t.Critical, ok = p[3].(bool)
		if !ok {
			return nil, fmt.Errorf("value for Critical was no bool")
		}
		fallthrough
	case l > 2:
		err := t.SetStart(p[2])
		if err != nil {
			return nil, err
		}
		fallthrough
	case l > 1:
		err := t.SetDuration(p[1])
		if err != nil {
			return nil, err
		}
		fallthrough
	case l > 0:
		t.Title, ok = p[0].(string)
		if !ok {
			return nil, fmt.Errorf("value for Title was no string")
		}
	}
	return t, nil
}

// CopyFields sets all (non-readonly) fields according to the given Task.
func (t *Task) CopyFields(task *Task) {
	if task != nil {
		t.Critical = task.Critical
		t.Active = task.Active
		t.Done = task.Done
		t.Title = task.Title
		// After should be copied as pointer to the same object
		t.After = task.After
		t.SetDuration(task)
		if task.Start == nil {
			t.Start = nil
		} else {
			timeNew := *task.Start
			t.Start = &timeNew
		}
	}
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
		tokens = append(tokens, t.id, t.Start.Format(time.RFC3339))
	} else if t.After != nil {
		tokens = append(tokens, t.id, "after "+t.After.id)
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
					`SetStart: "%s" is neither RFC3339 nor a valid Task ID`,
					tStart)
			}
			t.Start = &x
		}
	default:
		return fmt.Errorf("SetStart: unsupported type %T", start)
	}
	return nil
}

// Helperfunction to deduplicate code.
func (t *Task) setDurationFromTime(endTime *time.Time) (err error) {
	if t.Start != nil && endTime != nil {
		duration := endTime.Sub(*t.Start)
		t.Duration = &duration
		return nil
	}
	return fmt.Errorf("SetDuration: can't calculate duration from end time")
}

// Helperfunction to deduplicate code.
func (t *Task) setDurationFromTask(task *Task) {
	if task.Duration == nil {
		t.Duration = nil
	} else {
		newDur := *task.Duration
		t.Duration = &newDur
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
					`SetDuration: "%s" is neither a valid duration nor Task ID`,
					tDuration)
			}
			t.Duration = &x
		}
	default:
		return fmt.Errorf("SetDuration: unsupported type %T", duration)
	}
	return nil
}
