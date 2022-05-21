package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func InitDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	log.Println(psqlInfo)
	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln("Error Connecting to database", err)
	}

	Db = db
}
