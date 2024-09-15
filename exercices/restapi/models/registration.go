package models

import "course.go/restapi/db"

type Registration struct {
	ID      int64
	UserId  int64 `binding:"required"`
	EventId int64 `binding:"required"`
}

func (r *Registration) Save() error {
	insertRegistration := `
	INSERT INTO registrations(event_id, user_id) VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(insertRegistration)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(r.EventId, r.UserId)

	if err != nil {
		return err
	}

	r.ID, err = result.LastInsertId()
	return err
}

func (r Registration) DeleteRegistration() error {
	query := "DELETE FROM registration WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.ID)

	return err
}

func GetRegistrationById(id int64) (*Registration, error) {
	query := "SELECT * FROM registration WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return &Registration{}, err
	}

	row := stmt.QueryRow(id)

	var r Registration

	err = row.Scan(&r.ID, &r.UserId, &r.EventId)

	return &r, err

}
