package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	needInit := err != nil

	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if needInit {
		if err := createSchema(); err != nil {
			return err
		}
	}

	return nil
}

func createSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT '',
		title VARCHAR(255) NOT NULL DEFAULT '',
		comment TEXT NOT NULL DEFAULT '',
		repeat VARCHAR(128) NOT NULL DEFAULT ''
	);

	CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
	`

	_, err := DB.Exec(schema)
	return err
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// AddTask добавляет новую задачу в БД
func AddTask(task Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
