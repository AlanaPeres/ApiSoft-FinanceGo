package rotas

import (
	"ApiSoft-Finance/src/controllers"
	"net/http"
)

var rotaConta = Rota{
	URI:                "/contas/{cpf}",
	Metodo:             http.MethodGet,
	Funcao:             controllers.BuscarContaPorCpf,
	RequerAutenticacao: true,
}
