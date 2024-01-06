package banco

import (
	"ApiSoft-Finance/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o banco de dados retornar o objeto *sql.DB representando essa conexão.
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
