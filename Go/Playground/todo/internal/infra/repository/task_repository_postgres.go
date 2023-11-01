package repository

import (
	"context"
	"database/sql"

	"github.com/viniciusps01/internal/entity"
	"github.com/viniciusps01/pkg/apperrors"
)

type TaskRepositoryPostgres struct {
	Conn *sql.Conn
}

func NewTaskRepositoryPostgres(conn *sql.Conn) *TaskRepositoryPostgres {
	r := &TaskRepositoryPostgres{
		Conn: conn,
	}

	return r
}

func (r *TaskRepositoryPostgres) Create(t entity.Task) (int64, error) {
	var id int64

	err := r.Conn.QueryRowContext(context.TODO(), "insert into tasks(description, done) values ($1, $2) returning id",
		t.Description,
		t.Done,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskRepositoryPostgres) Read(ID int64) (*entity.Task, error) {
	res := r.Conn.QueryRowContext(context.TODO(), "select * from tasks where id=$1", ID)

	if err := res.Err(); err != nil {
		return nil, err
	}

	task := &entity.Task{}

	err := res.Scan(&task.ID, &task.Description, &task.Done)

	if err != nil {
		err := apperrors.NotFoundError{
			Message: "task not found",
		}

		return nil, err
	}

	return task, nil
}

func (r *TaskRepositoryPostgres) Update(t entity.Task) error {
	_, err := r.Conn.ExecContext(context.TODO(), "UPDATE tasks SET description=$1, done=$2 WHERE id=$3",
		t.Description,
		t.Done,
		t.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepositoryPostgres) Delete(ID int64) error {
	res := r.Conn.QueryRowContext(context.TODO(), "DELETE FROM tasks WHERE id=$1 returning id", ID)

	if err := res.Err(); err != nil {
		return err
	}

	var deletedID int64

	res.Scan(&deletedID)

	if deletedID != ID {
		err := apperrors.NotFoundError{
			Message: "task not found",
		}

		return err
	}

	return nil
}

func (r *TaskRepositoryPostgres) ReadAll() ([]*entity.Task, error) {
	res, err := r.Conn.QueryContext(context.TODO(), "select * from tasks")

	if err != nil {
		return nil, err
	}

	defer res.Close()

	tasks := []*entity.Task{}

	for res.Next() {
		var task entity.Task
		err := res.Scan(&task.ID, &task.Description, &task.Done)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}
