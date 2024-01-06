package models

import (
	"ApiSoft-Finance/src/seguranca"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	Cpf   uint64 `json:"cpf,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Cpf == 0 {
		return errors.New("O cpf é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}

	return nil
}
