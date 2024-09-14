package models

import (
	"errors"

	"course.go/restapi/db"
	"course.go/restapi/utils"
)

type Users struct {
	Id       int64
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

	u.Id = userID

	return err
}

func (u Users) ValidatedCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retreivedPwd string
	err := row.Scan(&u.Id, &retreivedPwd)

	if err != nil {
		return errors.New("credentials invalid")
	}

	if !(utils.ComparePwdHash(u.Password, retreivedPwd)) {
		return errors.New("credentials invalid")
	}

	return nil

}
