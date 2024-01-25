package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	-"github.com/mattn/go-sqlite3"
)

type Ticket struct {

	ID 	  int
	CreatedAt time.Time
	UserDate  time.Time
	PersonHelped string
	Issue        string
}

func main(){
	db, err := sql.open("sqlite3", "./tickets.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if it doesnt exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tickets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_date TIMESTAMP,
			person_helped TEXT,
			issue TEXT
		);
	`

	_. err = db.Exec(createTableQuery)
	if err != nil{
		log.Fatal(err)
	}

	insertTicketQuery := `
		INSERT INTO tickets (user_date, person_helped, issue) VALUES (?, ?, ?);
	`
	
	result, err := db.Exec(insertTicketQuery, time.Now(), "John Doe", "Printer Connection")
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted ticket with ID")

