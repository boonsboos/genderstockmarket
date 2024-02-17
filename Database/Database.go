package database

import (
	"context"
	util "gns500/Util"
	"log"

	"github.com/jackc/pgx/v4"
)

var DatabaseConnection Database = NewDatabase()

// simple holder for the database handle
//
// can build some DAO stuff around it
type Database struct {
	Conn *pgx.Conn
}

func NewDatabase() Database {
	conn, err := pgx.Connect(context.TODO(), util.Options.DatabaseURL+"/"+util.Options.DatabaseName+"?sslmode=disable")
	if err != nil {
		log.Println("Failed to connect to database:", err.Error())
	}
	log.Println("Database connection OK")

	return Database{
		Conn: conn,
	}
}