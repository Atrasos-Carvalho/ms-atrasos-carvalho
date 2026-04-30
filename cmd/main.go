package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	httphandler "ms-atrasos-carvalho/internal/delivery/http"
	"ms-atrasos-carvalho/internal/infra/postgres"
	"ms-atrasos-carvalho/internal/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("erro ao carregar .env:", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("DATABASE_URL não definida")
	}

	log.Println("conectando em:", dbURL)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("banco inacessível:", err)
	}

	registroRepo := postgres.NewRegistroRepo(db)
	registroUC   := usecase.NewRegistroUseCase(registroRepo)
	registroH    := httphandler.NewRegistroHandler(registroUC)

	router := httphandler.NewRouter(registroH)

	log.Println("servidor rodando em :8080")
	log.Fatal(router.Run(":8080"))
}