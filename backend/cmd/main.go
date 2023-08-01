package main

import (
	"log"
	"net/http"

	"github.com/nuradd1n/backend/db"
	"github.com/nuradd1n/backend/route"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed: %v\n", err)
	}

	mux := route.Router()
	log.Println("Запуск веб-сервера на http://localhost:8089/")
	err = http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Fatalf("ошибка при завершении работы сервера: %v\n", err)
	}
}
