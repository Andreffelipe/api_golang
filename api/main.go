package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*
instancia conexão com banco de dados
db, err: = sql.Open ("mysql", "<usuário>: <senha> @tcp (127.0.0.1:3306) / <nome do banco de dados>")
*/
var db *sql.DB
var err error

func main() {
	db, err := sql.Open("mysql", "<user>:<pass>@tcp(127.0.0.1:3306)/<database name>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	http.ListenAndServe(":8000", router)

}
