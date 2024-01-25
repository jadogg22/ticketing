package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Ticket represents the structure of a ticket.
type Ticket struct {
	ID           int
	CreatedAt    time.Time
	UserDate     time.Time
	PersonHelped string
	Issue        string
}

// Note represents the structure of a note.
type Note struct {
	ID        int
	TicketID  int
	CreatedAt time.Time
	Text      string
}

func createTicket(db *sql.DB, user_date time.Time, person_helped string, issue string) int {
	fmt.Printf("adding ticket to the db")

	insertTicketQuery := `
		INSERT INTO tickets (user_date, person_helped, issue) VALUES (?, ?, ?);
	`
	result, err := db.Exec(insertTicketQuery, user_date, person_helped, issue)
	if err != nil {
		log.Fatal(err)
	}

	// Get the ID of the inserted ticket
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted ticket with ID: %d\n", lastInsertID)
	return int(lastInsertID)
}

func createNote(db *sql.DB, ticket int, note string) {
	fmt.Printf("Adding Note to ticket: %d", ticket)

	// Example: Insert a note associated with the ticket
	insertNoteQuery := `
		INSERT INTO notes (ticket_id, text) VALUES (?, ?);
	`
	_, err := db.Exec(insertNoteQuery, ticket, note)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a note associated with the ticket.")

}

func getTicket(db *sql.DB, ticketID int) (*Ticket, error) {
	//querty for getting the specified ticket
	getTicketQuery := `
		SELECT id, created_at, user_date, person_helped, issue
		FROM tickets
		WHERE id = ?;

	`

	row := db.QueryRow(getTicketQuery, ticketID)

	var ticket Ticket
	err := row.Scan(&ticketID, &ticket.CreatedAt, &ticket.UserDate, &ticket.PersonHelped, &ticket.Issue)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Ticket with ID %d not found", ticketID)
		}
		return nil, err
	}
	return &ticket, nil

}

func main() {
	// Open or create the SQLite database file
	db, err := sql.Open("sqlite3", "./tickets.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the "tickets" table if it doesn't exist
	createTicketsTableQuery := `
		CREATE TABLE IF NOT EXISTS tickets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_date TIMESTAMP,
			person_helped TEXT,
			issue TEXT
		);
	`
	_, err = db.Exec(createTicketsTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "notes" table if it doesn't exist
	createNotesTableQuery := `
		CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ticket_id INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			text TEXT,
			FOREIGN KEY (ticket_id) REFERENCES tickets (id) ON DELETE CASCADE
		);
	`
	_, err = db.Exec(createNotesTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	ticket, err := getTicket(db, 2)

	fmt.Printf("Ticket ID: %d\n", ticket.ID)
	fmt.Printf("Created At: %s\n", ticket.CreatedAt)
	fmt.Printf("User Date: %s\n", ticket.UserDate)
	fmt.Printf("Person Helped: %s\n", ticket.PersonHelped)
	fmt.Printf("Issue: %s\n", ticket.Issue)
}
