package data_source

import (
	"context"
	"database/sql"

	"github.com/viniciusps01/internal/feature/task/entity"
	"github.com/viniciusps01/pkg/apperrors"
)

const (
	defaultLimit  = 10
	defaultOffset = 0
)

type TaskDataSourcePostgres struct {
	conn *sql.Conn
}

func NewTaskDataSourcePostgres(conn *sql.Conn) *TaskDataSourcePostgres {
	r := &TaskDataSourcePostgres{
		conn: conn,
	}

	return r
}

func (r *TaskDataSourcePostgres) Create(t entity.Task) (int64, error) {
	var id int64

	q := "INSERT INTO task(user_id, title, description, done) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.conn.QueryRowContext(context.TODO(), q,
		t.UserID,
		t.Title,
		t.Description,
		t.Done,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskDataSourcePostgres) Read(ID int64, userID string) (*entity.Task, error) {
	q := "SELECT id, user_id, title, description, done FROM task WHERE id=$1 AND user_id=$2"
	res := r.conn.QueryRowContext(context.TODO(), q, ID, userID)

	if err := res.Err(); err != nil {
		return nil, err
	}

	task := &entity.Task{}

	err := res.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Done)

	if err != nil {
		err := apperrors.NotFoundError{
			Message: "task not found",
		}

		return nil, err
	}

	return task, nil
}

func (r *TaskDataSourcePostgres) Update(t entity.Task) error {
	q := "UPDATE task SET title=$1, description=$2, done=$3 WHERE id=$4 AND user_id=$5"
	_, err := r.conn.ExecContext(context.TODO(), q,
		t.Title,
		t.Description,
		t.Done,
		t.ID,
		t.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskDataSourcePostgres) Delete(ID int64, userID string) error {
	q := "DELETE FROM task WHERE id=$1 AND user_id=$2 returning id"
	res := r.conn.QueryRowContext(context.TODO(), q, ID, userID)

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

func (r *TaskDataSourcePostgres) ReadAll(userID string, limit, offset *int) ([]*entity.Task, error) {
	if limit == nil {
		l := defaultLimit
		limit = &l
	}

	if offset == nil {
		o := defaultOffset
		offset = &o
	}

	q := "SELECT id, user_id, title, description, done FROM task WHERE user_id=$1 LIMIT $2 OFFSET $3"
	res, err := r.conn.QueryContext(context.TODO(), q, userID, *limit, *offset)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	tasks := []*entity.Task{}

	for res.Next() {
		var task entity.Task
		err := res.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Done)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}
