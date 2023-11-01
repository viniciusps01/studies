package entity

type Task struct {
	ID          int64
	Description string
	Done        bool
}

func NewTask(description string, done bool) *Task {
	task := &Task{
		Description: description,
		Done:        done,
	}

	return task
}
