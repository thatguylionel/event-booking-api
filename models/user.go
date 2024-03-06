package models

import (
	"tgl/eventapi/db"
	"tgl/eventapi/utils"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `binding:"required"   json:"email"`
	Password string `binding:"required"  json:"password"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()

	u.ID = int(userId)
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE lower(email) = lower(?)`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	return utils.ComparePasswords(retrievedPassword, u.Password)
}
