package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxIdleDbConnection = 5
	maxOpenDbConnection = 10
	maxDbLifetime       = 5 * time.Minute
)

func ConnectSql(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)

	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConnection)
	d.SetMaxIdleConns(maxIdleDbConnection)
	d.SetConnMaxLifetime(maxDbLifetime)

	err = testDb(d)

	if err != nil {
		return nil, err
	}

	dbConn.SQL = d

	return dbConn, nil
}

func testDb(db *sql.DB) error {
	err := db.Ping()

	if err != nil {
		return err
	}

	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
