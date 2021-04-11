package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/commons"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/lib/pq"
)

func main() {
	commons.InitDB()
	db, err := sql.Open("postgres", commons.CONNECTION_STRING)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})
}
