package repositorios

import (
	"ApiSoft-Finance/src/models"
	"database/sql"
)

type Transferencia struct {
	db *sql.DB
}

func NovoRepositorioDeTransferencia(db *sql.DB) *Transferencia {
	return &Transferencia{db}
}

func (repositorio *Transferencia) EncontrarContaPeloCPF(cpf uint64) (*models.ContaBancaria, error) {

	// Buscar a conta bancária pelo cpf do usuário logado
	var conta models.ContaBancaria
	err := repositorio.db.QueryRow("SELECT ContaBancariaId, Nome, Saldo FROM ContaBancaria WHERE Cpf = ?", cpf).Scan(&conta.ContaBancariaId, &conta.Nome, &conta.Saldo)
	if err != nil {
		return nil, err
	}

	return &conta, nil
}
func (repositorio *Transferencia) EncontrarContaPeloId(contaId uint64) (*models.ContaBancaria, error) {

	// Buscar a conta bancária pelo cpf do usuário logado
	var contaDestino models.ContaBancaria
	err := repositorio.db.QueryRow("SELECT ContaBancariaId, Nome, Saldo FROM ContaBancaria WHERE contaBancaria = ?", contaId).Scan(&contaDestino.ContaBancariaId, &contaDestino.Nome, &contaDestino.Saldo)
	if err != nil {
		return nil, err
	}

	return &contaDestino, nil
}

func (repositorio Transferencia) Transferir(transferencia models.Transferencia) (uint64, error) {
	// Buscar a conta vinculada ao cpf logado
	contaOrigem, err := repositorio.EncontrarContaPeloCPF(transferencia.ContaOrigemID)
	if err != nil {
		return 0, err
	}
	contaDestino, erro := repositorio.EncontrarContaPeloId(transferencia.ContaDestinoID)
	if contaOrigem == nil || contaDestino == nil {
		return 0, erro
	}

	// Validar saldo suficiente na conta de origem
	if contaOrigem.Saldo < transferencia.Valor {
		return 0, erro
	}
	// Atualizar o saldo da conta de destino
	contaOrigem.Saldo -= transferencia.Valor
	contaDestino.Saldo += transferencia.Valor

	statement, erro := repositorio.db.Prepare(
		"INSERT INTO Transacoes (tipo, valor, conta_origem_id, conta_destino_id ) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	// Executar a inserção na tabela Transacoes
	resultado, erro := statement.Exec("transferencia", transferencia.Valor, contaOrigem.ContaBancariaId, contaDestino.ContaBancariaId)
	if erro != nil {
		return 0, erro
	}

	statement, erro = repositorio.db.Prepare("UPDATE ContaBancaria SET Saldo = ? WHERE ContaBancariaId = ?", transferencia.Valor, contaOrigem.ContaBancariaId)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	// Atualizar o saldo da conta bancária no banco de dados
	
	// Obter o ID da última transferência inserida
	ultimoID, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoID), nil

}
