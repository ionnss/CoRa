package models

import (
	"database/sql"
	"fmt"
)

// Paciente representa os dados de um paciente
type Paciente struct {
	ID       int
	Nome     string
	CPF      string
	Email    string
	Telefone string
}

// PadronizarTelefone adiciona o prefixo 55 ao telefone, se necessário
func (p *Paciente) PadronizarTelefone() {
	if len(p.Telefone) > 0 && p.Telefone[:2] != "55" {
		p.Telefone = "55" + p.Telefone
	}
}

// InserirPaciente insere um novo paciente no banco de dados
func InserirPaciente(db *sql.DB, paciente Paciente) error {
	// Padroniza o telefone antes de inserir no banco
	paciente.PadronizarTelefone()

	query := `
		INSERT INTO pacientes (nome, cpf, email, telefone)
		VALUES (?, ?, ?, ?)
	`
	_, err := db.Exec(query, paciente.Nome, paciente.CPF, paciente.Email, paciente.Telefone)
	if err != nil {
		return fmt.Errorf("erro ao inserir paciente: %w", err)
	}
	return nil
}
