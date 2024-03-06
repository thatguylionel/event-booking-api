package models

import (
	"fmt"
	"tgl/eventapi/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"datetime"`
	UserID      int       `json:"userid"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	// protetects agains sql injection
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err

	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	fmt.Print("rows: ", rows)
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		fmt.Println(event)
		events = append(events, event)
	}

	return events, nil
}
