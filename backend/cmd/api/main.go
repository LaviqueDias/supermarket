package main

import (
	"log"
	"net/http"
	"github.com/LaviqueDias/supermarket/routes"
	"github.com/LaviqueDias/supermarket/internal/database"
	"github.com/LaviqueDias/supermarket/pkg/config"
)

func main() {
	db := database.Connect()
	defer db.Close()

	cfg := config.Get()

	router := routes.SetupRouter(db)

	log.Println("Servidor rodando na porta :" + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(cfg.ServerPort, router))
}