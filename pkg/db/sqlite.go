package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLite(dbPath string) (*sql.DB, error) {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("database %s does not exist", dbPath)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping SQLite database: %w", err)
	}

	log.Printf("Connected to SQLite database at %s", dbPath)
	return db, nil
}
