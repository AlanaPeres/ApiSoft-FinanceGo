package repositorios

import (
	"ApiSoft-Finance/src/models"
	"database/sql"
	"fmt"
)

// Representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um novo usuário no banco de dados
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (cpf, nome, email,senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	_, erro = statement.Exec(usuario.Cpf, usuario.Nome, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	return usuario.Cpf, nil
}

// Buscar traz todos os usuários que estão armazenados no banco de dados
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"SELECT cpf, nome, email FROM usuarios")
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()
	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.Cpf,
			&usuario.Nome,
			&usuario.Email,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorCpf(Cpf int64) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select cpf, nome, email from usuarios where cpf = ?",
		Cpf,
	)

	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Cpf,
			&usuario.Nome,
			&usuario.Email,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(Cpf uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, cpf = ?, email = ? where cpf = ?",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Cpf, usuario.Email, Cpf); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados.
func (repositorio Usuarios) Deletar(Cpf uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where cpf = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Cpf); erro != nil {
		return erro
	}
	return nil
}
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("select cpf, senha from usuarios where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Cpf, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(Cpf uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where cpf = ?", Cpf)
	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

// Atualizar senha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(Cpf uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where cpf = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, Cpf); erro != nil {
		return erro
	}

	return nil
}
