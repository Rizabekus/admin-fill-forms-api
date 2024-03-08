package app

import (
	"Rizabekus/admin-fill-forms-api/internal/handlers"
	"Rizabekus/admin-fill-forms-api/internal/services"
	"Rizabekus/admin-fill-forms-api/internal/storage"
	"Rizabekus/admin-fill-forms-api/pkg/loggers"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	loggers.InfoLog.Println("Loaded the configuration data from .env")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	err = storage.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}
	loggers.InfoLog.Println("Successfully connected to database")
	storage := storage.StorageInstance(db)
	service := services.ServiceInstance(storage)
	handler := handlers.HandlersInstance(service)

	Routes(handler)

	defer db.Close()
}
