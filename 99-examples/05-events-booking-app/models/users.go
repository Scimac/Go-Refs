package models

import (
	"errors"
	"events-booking/db"
	"events-booking/utils"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashNewPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.Id = id

	return nil
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var fetchedUser User
	err := row.Scan(&fetchedUser.Id, &fetchedUser.Password)
	if err != nil {
		return err
	}

	if !utils.CheckValidHashPassword(user.Password, fetchedUser.Password) {
		return errors.New("invalid creds")
	}

	user.Id = fetchedUser.Id
	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT id, email, password FROM users"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usersList []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, u)
	}

	return usersList, nil
}
