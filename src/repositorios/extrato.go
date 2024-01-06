package repositorios

import (
	"ApiSoft-Finance/src/models"
	"database/sql"
)

type Extrato struct {
	db *sql.DB
}

func (repositorio Extrato) Criar(contaBancaria models.Extrato) error {
	return nil
}

func (repositorio Extrato) Alterar(contaBancaria models.Extrato) error {
	return nil
}

func (repositorio Extrato) Deletar(id int) error {
	return nil
}
