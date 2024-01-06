package models

import "time"

type Transferencia struct {
	ID                    uint64    `json:"id"`
	Tipo                  string    `json:"tipo"`
	Valor                 float64   `json:"valor,omitempty"`
	DataHoraTransferencia time.Time `json:"data_hora,omitempty"`
	Cpf                   uint64    `json:"cpf"`
	ContaOrigemID         uint64    `json:"conta_origem_id"`
	ContaDestinoID        uint64    `json:"conta_destino_id"`
}
