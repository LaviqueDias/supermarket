package database

import (
	"database/sql"
	"log"
	"github.com/LaviqueDias/supermarket/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	cfg := config.Get()
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao MySQL:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao testar conexão:", err)
	}

	log.Println("✓ Conectado ao MySQL")
	return db
}


