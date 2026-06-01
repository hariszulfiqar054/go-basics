package models

import (
	"example.com/rest_api/db"
	"example.com/rest_api/utils"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CreatedAt string `json:"created_at"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return err
}

func GetUserByEmail(email string) (*User, error) {
	query := `SELECT * FROM users WHERE email = ?`
	row, err := db.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var user User
	if row.Next() {
		err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}
