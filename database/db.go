package database

import (
	"database/sql"
	"fmt"
	"log"

	"public-vault-ms/config"
)

var DB *sql.DB

func InitDatabase(cfg *config.Config) {
	var err error

	DB, err = sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("No se puede establecer conexión con la base de datos: %v", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL")

	// Crea la tabla si no existe
	createTable := `
	CREATE TABLE IF NOT EXISTS cards (
		token TEXT PRIMARY KEY,
		encrypted_card TEXT NOT NULL
	);
	`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Error al crear la tabla: %v", err)
	}
}
