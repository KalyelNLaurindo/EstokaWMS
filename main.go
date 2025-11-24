package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Conectar ao banco SQLite
	db, err := sql.Open("sqlite3", "./data/estoka_wms.db")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	log.Println("Conexão com o banco de dados SQLite estabelecida com sucesso!")
}
