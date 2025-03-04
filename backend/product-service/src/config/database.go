package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Erreur de connexion à la base de données :", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("❌ Impossible de ping PostgreSQL :", err)
	}

	fmt.Println("✅ Connexion à PostgreSQL réussie")
	DB = db
}
