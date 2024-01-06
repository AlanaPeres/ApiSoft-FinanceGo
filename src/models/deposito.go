package models

import "time"

type Deposito struct {
	ID              uint64    `json:"id"`
	Valor           float64   `json:"valor,omitempty"`
	DataHora        time.Time `json:"data_hora,omitempty"`
	Cpf             uint64    `json:"cpf"`
	ContaBancariaId uint64    `json:"conta_bancaria_id"`
}
