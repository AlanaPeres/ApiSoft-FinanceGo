package rotas

import (
	"ApiSoft-Finance/src/controllers"
	"net/http"
)

var rotaTransferencia = Rota{
	URI:                "/transferencia/{cpf}",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Depositar,
	RequerAutenticacao: false,
}
