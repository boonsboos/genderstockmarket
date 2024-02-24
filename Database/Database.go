package database

import (
	"context"
	"log"
	util "spectrum300/Util"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool = NewDatabase()

func NewDatabase() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), util.Options.DatabaseURL+"/"+util.Options.DatabaseName+"?sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err.Error())
	}
	log.Println("Database connection OK")

	return conn
}
