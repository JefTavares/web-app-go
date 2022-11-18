package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// retorna uma conexão com o banco de dados (retorno *sql.DB)
// Para exportar a função para outros package a primeira letra deve ser maiuscula
func ConectaComBancoDeDados() *sql.DB {
	//teste - a abertura sera feita nas proprias função que faz a chamadas
	//db := conectaComBancoDeDados() /*Abre a conexão com o banco */
	//defer db.Close()
	conexao := "user=jeftavares dbname=loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
