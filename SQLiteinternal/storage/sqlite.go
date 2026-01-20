package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "favorites.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS favorites (
		id TEXT PRIMARY KEY,
		name TEXT,
		blank TEXT
	);
	`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
