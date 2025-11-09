package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)


func InitDB()*sql.DB{
	godotenv.Load(".env")
	connStr := os.Getenv("DATABASE_URL")
	
	db , err := sql.Open("pgx" , connStr)
	if err != nil{
		log.Fatalf("Erro ao abrir conexão: %v" , err)
	}

	err = db.Ping()
	if err != nil {
        log.Fatalf("Erro ao conectar ao banco: %v", err)
    }

	return db
}