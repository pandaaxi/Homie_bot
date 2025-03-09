package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB opens the SQLite database and executes necessary migrations.
func InitDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS vps (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			remark TEXT,
			ip TEXT,
			region TEXT,
			provider TEXT,
			price REAL,
			total_data INTEGER,
			reset_times INTEGER,
			available_data INTEGER,
			reset_date TEXT,
			calculation_type TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS usage_logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			vps_id INTEGER,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			in_usage INTEGER,
			out_usage INTEGER,
			total_usage INTEGER,
			FOREIGN KEY(vps_id) REFERENCES vps(id)
		);`,
		`CREATE TABLE IF NOT EXISTS alerts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			vps_id INTEGER,
			threshold REAL,
			alert_type TEXT,
			FOREIGN KEY(vps_id) REFERENCES vps(id)
		);`,
	}

	for _, q := range queries {
		_, err = DB.Exec(q)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			return err
		}
	}
	return nil
}
