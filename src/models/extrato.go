package models

import "time"

// Estrutura de dados para representar um extrato bancário
type Extrato struct {
	Id             int
	Datetime       time.Time
	Descricao      string
	ValorTransacao float64
	SaldoAnterior  float64
	SaldoAtual     float64
}
