package controllers

import (
	"ApiSoft-Finance/src/autenticacao"
	"ApiSoft-Finance/src/banco"
	"ApiSoft-Finance/src/models"
	"ApiSoft-Finance/src/repositorios"
	"ApiSoft-Finance/src/respostas"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Depositar(w http.ResponseWriter, r *http.Request) {
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

	var deposito models.Deposito
	if erro = json.Unmarshal(corpoRequisicao, &deposito); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	//Validar o depósito
	if deposito.Valor <= 0 {
		respostas.Erro(w, http.StatusBadRequest, errors.New("O valor do depósito deve ser maior que 0"))
		return
	}
	// Definir ContaBancariaID como usuárioID
	deposito.ContaBancariaId = usuarioID
	fmt.Println(deposito.ContaBancariaId)

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeDeposito(db)
	deposito.ID, erro = repositorio.Depositar(deposito)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, deposito)

}
