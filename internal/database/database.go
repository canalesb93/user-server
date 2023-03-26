package database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

func New() (*Database, error) {
	dbPath := filepath.Join(".", "users-test.db")
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &Database{Conn: conn}, nil
}

func (d *Database) Close() error {
	err := d.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}