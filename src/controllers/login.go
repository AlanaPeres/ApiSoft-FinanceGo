package controllers

import (
	"ApiSoft-Finance/src/autenticacao"
	"ApiSoft-Finance/src/banco"
	"ApiSoft-Finance/src/models"
	"ApiSoft-Finance/src/repositorios"
	"ApiSoft-Finance/src/respostas"
	"ApiSoft-Finance/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return

	}

	db, erro := banco.Conectar()
	if erro != nil {
		//erro no servidor
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorios := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioDoBanco, erro := repositorios.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioDoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioDoBanco.Cpf)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Write([]byte(token))
}
