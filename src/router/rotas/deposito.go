package rotas

import (
	"ApiSoft-Finance/src/controllers"
	"net/http"
)

var rotaDeposito = Rota{
	URI:                "/deposito/{cpf}",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Depositar,
	RequerAutenticacao: false,
}
