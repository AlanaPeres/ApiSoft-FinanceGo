package repositorios

import (
	"ApiSoft-Finance/src/models"
	"database/sql"
)

type ContaBancaria struct {
	db *sql.DB
}

// NovoRepositoriDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDeConta(db *sql.DB) *ContaBancaria {
	return &ContaBancaria{db}
}

// Criar insere uma conta no banco de dados
func (repositorio ContaBancaria) Criar(contaBancaria models.ContaBancaria) (uint64, error) {
	var ultimoID sql.NullInt64
	erro := repositorio.db.QueryRow("SELECT MAX(ContaBancariaId) FROM ContaBancaria").Scan(&ultimoID)
	if erro != nil {
		if erro == sql.ErrNoRows {
			// Não há nenhuma conta bancária no banco de dados, então definimos um valor inicial.
			contaBancaria.ContaBancariaId = 1 // ou qualquer outro valor inicial desejado
		} else {
			return 0, erro
		}
	} else {
		if ultimoID.Valid {
			contaBancaria.ContaBancariaId = uint64(ultimoID.Int64) + 1
		} else {
			contaBancaria.ContaBancariaId = 1
		}
	}

	// Inserir a nova conta no banco de dados
	statement, erro := repositorio.db.Prepare("INSERT INTO ContaBancaria (Agencia, ContaBancariaId, Nome, Cpf, Saldo) VALUES (?, ?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()
	resultado, erro := statement.Exec("1515", contaBancaria.ContaBancariaId, contaBancaria.Nome, contaBancaria.Cpf, 0.0) // Saldo inicial 0.0
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que estão armazenados no banco de dados
func (repositorio ContaBancaria) BuscarContaPorCpf(Cpf int64) (models.ContaBancaria, error) {

	linhas, erro := repositorio.db.Query(
		"select agencia, contaBancariaId, nome, cpf, saldo from contaBancaria where cpf = ?",
		Cpf,
	)

	if erro != nil {
		return models.ContaBancaria{}, erro
	}
	defer linhas.Close()

	var conta models.ContaBancaria

	if linhas.Next() {
		if erro = linhas.Scan(
			&conta.Agencia,
			&conta.ContaBancariaId,
			&conta.Nome,
			&conta.Cpf,
			&conta.Saldo,
		); erro != nil {
			return models.ContaBancaria{}, erro
		}

	}
	return conta, nil
}

func (r *ContaBancaria) AtualizarSaldo(id uint64, novoSaldo float64) error {
	_, err := r.db.Exec("UPDATE ContaBancaria SET saldo = ? WHERE contaBancariaId = ?", novoSaldo, id)
	return err
}
