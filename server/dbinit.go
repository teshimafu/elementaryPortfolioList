package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gocraft/dbr/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
)

//Migration is called before server start
func Migration(filePath string) {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		HOST := os.Getenv("POSTGRES_HOST")
		PORT := os.Getenv("POSTGRES_PORT")
		USER := os.Getenv("POSTGRES_USER")
		PASS := os.Getenv("POSTGRES_PASSWORD")
		DBNAME := os.Getenv("POSTGRES_DBNAME")
		databaseURL = "postgres://" + USER + ":" + PASS + "@" + HOST + ":" + PORT + "/" + DBNAME + "?sslmode=disable"
	}
	m, err := migrate.New(
		filePath,
		databaseURL,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
	}
}

func gocraftConnection() (*dbr.Session, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		HOST := os.Getenv("POSTGRES_HOST")
		PORT := os.Getenv("POSTGRES_PORT")
		USER := os.Getenv("POSTGRES_USER")
		PASS := os.Getenv("POSTGRES_PASSWORD")
		DBNAME := os.Getenv("POSTGRES_DBNAME")
		databaseURL = "host=" + HOST + " port=" + PORT + " user=" + USER + " dbname=" + DBNAME + " password=" + PASS + " sslmode=disable"
	}
	conn, err := dbr.Open("postgres", databaseURL, nil)
	if err != nil {
		fmt.Printf("postgres error:%+v\n", err)
		return nil, err
	}
	conn.SetMaxOpenConns(10)
	return conn.NewSession(nil), nil
}
