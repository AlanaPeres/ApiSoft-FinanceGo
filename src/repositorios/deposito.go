package repositorios

import (
	"ApiSoft-Finance/src/models"
	"database/sql"
)

type Deposito struct {
	db *sql.DB
}

func NovoRepositorioDeDeposito(db *sql.DB) *Deposito {
	return &Deposito{db}
}

func (repositorio *Deposito) EncontrarContaPeloCPF(cpf uint64) (*models.ContaBancaria, error) {

	var conta models.ContaBancaria
	err := repositorio.db.QueryRow("SELECT ContaBancariaId, Nome, Saldo FROM ContaBancaria WHERE Cpf = ?", cpf).Scan(&conta.ContaBancariaId, &conta.Nome, &conta.Saldo)
	if err != nil {
		return nil, err
	}
	return &conta, nil
}

func (repositorio Deposito) Depositar(deposito models.Deposito) (uint64, error) {
	// Encontrar a conta bancária de origem usando o cpf da conta do usuário
	conta, erro := repositorio.EncontrarContaPeloCPF(deposito.ContaBancariaId)
	if erro != nil {
		return 0, erro
	}
	// Preparar o comando SQL para inserir na tabela Transacoes
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO Transacoes (tipo, valor, conta_origem_id ) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	// Executar a inserção na tabela Transacoes
	resultado, erro := statement.Exec("deposito", deposito.Valor, conta.ContaBancariaId)
	if erro != nil {
		return 0, erro
	}

	statement, erro = repositorio.db.Prepare("UPDATE ContaBancaria SET Saldo = ? WHERE ContaBancariaId = ?")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(conta.Saldo+deposito.Valor, conta.ContaBancariaId); erro != nil {
		return 0, erro
	}

	// Obter o ID do último depósito inserido
	ultimoID, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoID), nil

}
