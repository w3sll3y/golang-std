package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // CHAVE DUPLICADA
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	// COMO EH UMA TRANSACAO, NENHUM INSERT VAI SER INSERIDO
	// SE FOSSE DIRETO DO DB.PREPARE O UNICO QUE NAO SERIA INSERIDO SERIA O TIAGO
	// EM GO EH PRECISO TRATAR O ERRO, CASO CONTRARIO SERIA IGUAL O DB.PREPARE, E O UNICO QUE NAO SERIA INSERIDO SERIA O TIAGO
}
