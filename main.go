package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Vasilinateg/go-final-todo/pkg/api"
	"github.com/Vasilinateg/go-final-todo/pkg/db"
)

func main() {
	// Определяем порт
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	// Инициализация базы данных
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	if err := db.Init(dbFile); err != nil {
		log.Printf("Ошибка инициализации БД: %v", err)
		return
	}
	defer db.Close()

	// Инициализация API-обработчиков
	api.Init()

	// Директория с веб-файлами
	webDir := "./web"

	// Создаём файловый сервер
	fs := http.FileServer(http.Dir(webDir))

	// Обработчик для всех запросов
	http.Handle("/", fs)

	log.Printf("Сервер запущен на порту %s", port)
	log.Printf("Откройте http://localhost:%s/ в браузере", port)
	log.Printf("База данных: %s", dbFile)

	// Запускаем сервер
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Ошибка запуска сервера: %v", err)
	}
}
