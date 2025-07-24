package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"perfumestore/internal/db"
	"perfumestore/internal/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	database := db.Init()
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Миграция не удалась", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Сервер стартует на порту:", port)
	http.ListenAndServe(":"+port, nil)
}
