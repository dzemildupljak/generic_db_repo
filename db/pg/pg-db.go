package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dzemildupljak/web_app_with_unitest/config"
	_ "github.com/lib/pq"
)

var ErrNotFound = errors.New("database record not found\n")
var ErrIvalidFilters = errors.New("invalid filters\n")

var database *sql.DB

func DB() *sql.DB {
	return database
}

func Init() {
	c := config.Instance().Database
	dbString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.Name)

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open database connection: %w", err))
	} else {
		fmt.Println("successfully connected to database")
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("failed to ping database: %w", err))
	} else {
		fmt.Println("succeed to ping database")
	}
	database = db
}

func Close() {
	_ = database.Close()
}
