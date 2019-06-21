package gantt

import (
	"fmt"
)

type Section struct {
	id    string
	gantt *Gantt
	tasks []*Task
}

func (s *Section) String() (renderedElement string) {
	text := fmt.Sprintln("section", s.id)
	for _, task := range s.tasks {
		text += task.String()
	}
	return text
}

// AddTask ...
// If error
func (s *Section) AddTask(id string, init ...interface{}) (newTask *Task, err error) {
	_, alreadyExists := s.gantt.tasksMap[id]
	if alreadyExists {
		return nil, fmt.Errorf("id already exists")
	}
	t, e := taskNew(id, s.gantt, s, init)
	if e != nil {
		return nil, e
	}
	s.gantt.tasksMap[id] = t
	s.tasks = append(s.tasks, t)
	return t, nil
}

/*
   section A section
   Completed task            :done,    des1, 2014-01-06,2014-01-08
   Active task               :active,  des2, 2014-01-09, 3d
   Future task               :         des3, after des2, 5d
   Future task2              :         des4, after des3, 5d

   section Critical tasks
   Completed task in the critical line :crit, done, 2014-01-06,24h
   Implement parser and jison          :crit, done, after des1, 2d
   Create tests for parser             :crit, active, 3d
   Future task in critical line        :crit, 5d
   Create tests for renderer           :2d
   Add to mermaid                      :1d

   section Documentation
   Describe gantt syntax               :active, a1, after des1, 3d
   Add gantt diagram to demo page      :after a1  , 20h
   Add another diagram to demo page    :doc1, after a1  , 48h

   section Last section
   Describe gantt syntax               :after doc1, 3d
   Add gantt diagram to demo page      :20h
   Add another diagram to demo page    :48h
*/
