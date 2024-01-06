package controllers

import (
	"ApiSoft-Finance/src/banco"
	"ApiSoft-Finance/src/repositorios"
	"ApiSoft-Finance/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarContaPorCpf(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	cpf, erro := strconv.ParseUint(parametros["cpf"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorios := repositorios.NovoRepositorioDeConta(db)
	conta, erro := repositorios.BuscarContaPorCpf(int64(cpf))
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, conta)
}
