package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "modernc.org/sqlite" // Driver SQLite
)

// InitDB inicializa a conexão com o banco de dados
func InitDB() (*sql.DB, error) {
	// Abre a conexão com o banco de dados local
	db, err := sql.Open("sqlite", "./cora.db")
	if err != nil {
		return nil, fmt.Errorf("não foi possível abrir o banco de dados: %w", err)
	}

	// Carrega e aplica as instruções de criação de tabelas
	fmt.Println("Criando tabelas no banco de dados...")
	err = createTables(db)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar tabelas: %w", err)
	}

	fmt.Println("Tabelas criadas com sucesso!")
	return db, nil
}

// createTables lê o arquivo SQL cora/db/create_tables.sql e cria as tabelas no banco de dados
func createTables(db *sql.DB) error {
	// Lê o arquivo de criação de tabelas
	tablesSQL, err := ioutil.ReadFile("./db/create_tables.sql")
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo de criação de tabelas: %w", err)
	}

	// Executa o SQL para criar tabelas
	_, err = db.Exec(string(tablesSQL))
	if err != nil {
		return fmt.Errorf("erro ao executar as instruções de criação de tabelas: %w", err)
	}

	return nil
}

// main é o ponto de entrada da aplicação CoRa
func main() {
	// Inicializando banco de dados
	fmt.Println("Inicializando banco de dados...")
	db, err := InitDB()
	if err != nil {
		fmt.Printf("Erro ao inicializar o banco de dados: %v\n", err)
		return
	}
	defer db.Close()

	fmt.Println("Banco de dados inicializado com sucesso!")
}
