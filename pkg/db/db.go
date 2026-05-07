package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Init инициализирует базу данных
func Init(dbFile string) error {
	// Проверяем, существует ли файл БД
	_, err := os.Stat(dbFile)
	needInit := err != nil

	// Открываем БД
	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	// Если файла не было, создаём таблицу и индекс
	if needInit {
		if err := createSchema(); err != nil {
			DB.Close()
			return err
		}
	}

	return nil
}

// createSchema создаёт таблицу и индекс
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

// Close закрывает соединение с БД
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
