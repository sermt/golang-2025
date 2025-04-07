package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/util"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SaveUser(u User) error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	// asignar el ID generado a la instancia de usuario
	return nil
}

func ValidateCredentials(u User) (User, error) {
	query := `SELECT id, password FROM users WHERE email = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	var hashedPassword string
	var userID int
	err = stmt.QueryRow(u.Email).Scan(&userID, &hashedPassword)
	if err != nil {
		return User{}, err
	}

	if err := util.ComparePasswords(hashedPassword, u.Password); err != nil {
		return User{}, errors.New("invalid credentials")
	}

	// Devolver usuario con ID y email
	return User{
		ID:    userID,
		Email: u.Email,
	}, nil
}
