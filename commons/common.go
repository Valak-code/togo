package commons

import "fmt"

const (
	host     = "postgres-db"
	port     = 5432
	user     = "root"
	password = "password"
	dbname   = "togo"
)

var CONNECTION_STRING = ""

func InitDB() {
	CONNECTION_STRING = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
