package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Vasilinateg/go-final-todo/pkg/db"
)

func main() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	if err := db.Init(dbFile); err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}
	defer db.Close()

	webDir := "./web"

	fs := http.FileServer(http.Dir(webDir))

	http.Handle("/", fs)

	log.Printf("Сервер запущен на порту %s", port)
	log.Printf("Откройте http://localhost:%s/ в браузере", port)
	log.Printf("База данных: %s", dbFile)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
