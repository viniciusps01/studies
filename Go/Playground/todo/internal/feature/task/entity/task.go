package entity

type Task struct {
	ID          int64
	UserID      string
	Title       string
	Description string
	Done        bool
}

func NewTask(userID, title, description string, done bool) *Task {
	task := &Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Done:        done,
	}

	return task
}
