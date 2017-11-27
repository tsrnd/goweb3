package database

import (
	"log"
	"database/sql"
)
var (
	// SQL wrapper
	SQL *sql.DB
)
type congfigure interface {
	connect() (*sql.DB, error)
}
const (
	// TypeBolt is BoltDB
	TypeBolt string = "Bolt"
	// TypeMongoDB is MongoDB
	TypeMongoDB string = "MongoDB"
	// TypeMySQL is MySQL
	TypeMySQL string = "MySQL"
	TypePosgres string = "Postgres"
)
type Type string
type Connections struct {
	Postgres PostgresInfo
}
type Info struct {
	Driver string
	Connections Connections
}
func Connect(d Info) {
	var err error
	switch d.Driver {
	case TypePosgres:
		SQL, err = connect(d.Connections.Postgres)
		if err != nil {
			log.Println("SQL Driver Error", err)
		}
	default:
		log.Println("No registered database in config")
	}
}
func connect(g congfigure) (*sql.DB, error) { 
	return g.connect()
}