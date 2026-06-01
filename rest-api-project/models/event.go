package models

import (
	"time"

	"example.com/rest_api/db"
)

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (title, description, user_id) 
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(e.Title, e.Description, e.UserID)
	defer stmt.Close()
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.ID = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.CreatedAt, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}

func GetEventByID(id string) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var event Event
	if row.Next() {
		err := row.Scan(&event.ID, &event.Title, &event.Description, &event.CreatedAt, &event.UserID)
		if err != nil {
			return nil, err
		}
		return &event, nil
	}
	return nil, nil

}

func DeleteEventByID(id string) error {
	query := `DELETE FROM events WHERE id = ?`
	row, err := db.DB.Query(query, id)
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}
