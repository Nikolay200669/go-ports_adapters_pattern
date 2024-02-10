// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Nikolay200669/Ports_and_Adapters_patern/application"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Инициализация базы данных
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Создание экземпляра приложения
	app := application.NewApp(db)

	// Настройка роутера Gin
	router := app.SetupRouter()

	// Запуск HTTP-сервера
	port := 8080
	log.Printf("Starting server on :%d...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
