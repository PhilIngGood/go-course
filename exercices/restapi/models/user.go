package models

import (
	"course.go/restapi/db"
	"course.go/restapi/utils"
)

type Users struct {
	id       int64
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error {
	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?);"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPwd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Name, u.Email, hashedPwd)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	u.id = userID

	return err
}
