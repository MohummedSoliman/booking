// Package driver for set up the db driver configuration.
package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the db connection pool.
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDBConn = 10
	maxIdelDBConn = 5
	maxDBLifeTime = 5 * time.Second
)

// ConnectSQL creates database sql pool for Postgres.
func ConnectSQL(dbstring string) (*DB, error) {
	db, err := NewDB(dbstring)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdelDBConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	dbConn.SQL = db

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB try to ping to database
func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func NewDB(dbString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
