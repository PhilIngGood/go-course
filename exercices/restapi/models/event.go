package models

import (
	"time"

	"course.go/restapi/db"
)

type Event struct {
	ID          int64
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

// func NewEvent(id int, title, description, location string, date time.Time, userid int) Event {
// 	return Event{
// 		ID:          id,
// 		Title:       title,
// 		Description: description,
// 		Location:    location,
// 		DateTime:    date,
// 		UserID:      userid,
// 	}
// }

func (e Event) Save() error {
	insertEvent := "INSERT INTO events(name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?);"

	// Preparation of the insertion, can protect against sql injection
	stmt, err := db.DB.Prepare(insertEvent)

	if err != nil {
		return err
	}
	defer stmt.Close()
	// Inserting data
	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	// Grabing the ID returned by the "KEY AUTOINCREMENT"
	id, err := result.LastInsertId()
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
	for rows.Next() {
		var event Event

		// Scan works kinda like fmt.Scan, as it requieres pointer to variable and will store data into them. One pointer is requeried for every single properties
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil

}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?;"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
