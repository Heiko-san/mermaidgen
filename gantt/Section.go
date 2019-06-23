package gantt

import (
	"fmt"
)

// Section represents gantt sections that can be added to the Gantt diagram.
// Create an instance of Section via Gantt's AddSection method, do not create
// instances directly. Already defined IDs can be looked up via Gantt's
// GetSection method or iterated over via its ListSections method.
type Section struct {
	id    string
	gantt *Gantt
	tasks []*Task
}

// Private constructor for use in Add-functions.
func sectionNew(i string, g *Gantt) (*Section, error) {
	_, alreadyExists := g.sectionsMap[i]
	if alreadyExists {
		return nil, fmt.Errorf("id already exists")
	}
	return &Section{id: i, gantt: g}, nil
}

// ID provides access to the Sections readonly field id. This is used as the
// Section's title since the title needs to be unique. However the strict rules
// for Task's IDs don't apply here, feel free to use spaces and special
// characters.
func (s *Section) ID() (id string) {
	return s.id
}

// Gantt provides access to the top level Gantt diagram to be able to access
// Adder, Getter and Lister methods.
func (s *Section) Gantt() (topLevel *Gantt) {
	return s.gantt
}

// String renders this diagram element to a section definition line.
func (s *Section) String() (renderedElement string) {
	renderedElement = fmt.Sprintln("section", s.id)
	for _, task := range s.tasks {
		renderedElement += task.String()
	}
	return
}

// AddTask is used to add a new Task to this Section. If the provided ID already
// exists or is invalid, no new Task is created and an error is returned.
// The ID can later be used to look up the created Task using Gantt's GetTask
// method. Optional initializer parameters can be given in the order Title,
// Duration, Start, Critical, Active, Done. Duration and Start are set via
// Task's SetDuration and SetStart respectively.
func (s *Section) AddTask(id string, init ...interface{}) (newTask *Task, err error) {
	newTask, err = taskNew(id, s.gantt, s, init)
	if err != nil {
		return
	}
	s.gantt.tasksMap[id] = newTask
	s.tasks = append(s.tasks, newTask)
	return
}

// ListLocalTasks returns a slice of all Tasks previously added to this
// Section in the order they were defined.
func (s *Section) ListLocalTasks() (localTasks []*Task) {
	localTasks = make([]*Task, len(s.tasks))
	copy(localTasks, s.tasks)
	return
}
