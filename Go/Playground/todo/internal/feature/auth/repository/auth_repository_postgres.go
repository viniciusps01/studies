package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/viniciusps01/internal/feature/auth/entity"
	"github.com/viniciusps01/pkg/apperrors"
)

const (
	defaultLimit  = 10
	defaultOffset = 0
)

type AuthRepositoryPostgres struct {
	conn *sql.Conn
}

func NewAuthRepositoryPostgres(c *sql.Conn) AuthRepositoryPostgres {
	return AuthRepositoryPostgres{
		conn: c,
	}
}

func (r AuthRepositoryPostgres) Create(u entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	q := "INSERT INTO user_account(first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	row := r.conn.QueryRowContext(
		ctx,
		q,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	if err := row.Scan(&u.ID); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r AuthRepositoryPostgres) readUser(param, value string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	var user entity.User

	var permissionsJSON []byte

	q := fmt.Sprintf("SELECT u.id, u.email, u.first_name, u.last_name, u.password, json_agg(json_build_object('id', p.id, 'name', p.name)) as permissions from user_account u left join role_permission rp on u.role_id=rp.role_id left join permission p on p.id=rp.permission_id where u.%s=$1 group by u.id, u.email, p.id",
		param)
	err := r.conn.QueryRowContext(ctx, q, value).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&permissionsJSON,
	)

	if err != nil {
		return nil, apperrors.NotFoundError{
			Message: "User not found",
		}
	}

	if err := json.Unmarshal(permissionsJSON, &user.Permissions); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r AuthRepositoryPostgres) Read(ID string) (*entity.User, error) {
	return r.readUser("id", ID)
}

func (r AuthRepositoryPostgres) ReadUserByEmail(email string) (*entity.User, error) {
	return r.readUser("email", email)
}

func (r AuthRepositoryPostgres) Update(user entity.User) (*entity.User, error) {
	return nil, nil
}

func (r AuthRepositoryPostgres) Delete(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	var deletedID string

	q := "DELETE FROM user_account WHERE id=$1 RETURNING id"
	err := r.conn.QueryRowContext(ctx, q, ID).Scan(&deletedID)

	if err != nil {
		return err
	}

	if deletedID != ID {
		return apperrors.NotFoundError{
			Message: "User not found",
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (r AuthRepositoryPostgres) ReadAllUsers(limit, offset *int) (*[]entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if limit == nil {
		l := defaultLimit
		limit = &l
	}

	if offset == nil {
		o := defaultOffset
		offset = &o
	}

	q := "SELECT u.id, u.email, u.first_name, u.last_name, u.password, json_agg(json_build_object('id', p.id, 'name', p.name)) as permissions from user_account u left join role_permission rp on u.role_id=rp.role_id left join permission p on p.id=rp.permission_id group by u.id, u.email, p.id LIMIT $1 OFFSET $2"

	rows, err := r.conn.QueryContext(ctx, q, limit, offset)

	if err != nil {
		return nil, err
	}

	users := []entity.User{}

	for rows.Next() {

		var user entity.User
		var permissionsJSON []byte

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&permissionsJSON,
		)

		if err != nil {
			return nil, apperrors.InternalServerError{
				Message: "Failed to unmarshal user data",
			}
		}

		if err := json.Unmarshal(permissionsJSON, &user.Permissions); err != nil {
			return nil, apperrors.InternalServerError{
				Message: err.Error(),
			}
		}

		users = append(users, user)
	}

	return &users, nil
}
