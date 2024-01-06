package controllers

import (
	"ApiSoft-Finance/src/autenticacao"
	"ApiSoft-Finance/src/banco"
	"ApiSoft-Finance/src/models"
	"ApiSoft-Finance/src/repositorios"
	"ApiSoft-Finance/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Transferencia(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var transferencia models.Transferencia
	if erro = json.Unmarshal(corpoRequisicao, &transferencia); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	//Validar o depósito
	if transferencia.Valor <= 0 {
		respostas.Erro(w, http.StatusBadRequest, errors.New("O valor do depósito deve ser maior que 0"))
		return
	}
	// Definir ContaBancariaID como usuárioID
	transferencia.ContaOrigemID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeTransferencia(db)
	transferencia.ID, erro = repositorio.Transferir(transferencia)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, transferencia)

}
