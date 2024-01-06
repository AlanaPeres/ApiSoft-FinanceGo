package models

type ContaBancaria struct {
	Agencia         int     `json:"agencia,omitempty"`
	ContaBancariaId uint64  `json:"contaBancaria,omitempty"`
	Nome            string  `json:"nome,omitempty"`
	Cpf             uint64  `json:"cpf,omitempty"`
	Saldo           float64 `json:"saldo"`
}
