package db

import (
	"fmt"
	"log"
	"project/handlers"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	conn *sqlx.DB
}

func ConnectDB() *Connection {
	dbUser := handlers.DotEnvVariable("DB_USER")
	dbPassword := handlers.DotEnvVariable("DB_PASSWORD")
	dbHost := handlers.DotEnvVariable("DB_HOST")
	dbName := handlers.DotEnvVariable("DB_NAME")
	dbPort := handlers.DotEnvVariable("DB_PORT")

	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Open("mysql", uri)
	if err != nil {
		log.Printf("database error: %v\n", err.Error())
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Printf("database error: %v\n", "unable to connect database")
		log.Fatal("unable to connect database")
	}
	color.Green("⛁ Connected to Database")

	return &Connection{
		conn: db,
	}
}
