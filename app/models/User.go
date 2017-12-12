package models

import (
	"github.com/goweb3/app/shared/database"
	"github.com/goweb3/app/shared/passhash"
)

type User struct {
	BaseModel
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

/**
*
* Hash password of user
**/
func (user *User) HashPassword() error {
	pass, err := passhash.HashString(user.Password)
	if err == nil {
		user.Password = pass
	}
	return err
}

/**
*
* Find user by name
**/
func (user *User) FindByName(name string) (err error) {
	err = database.SQL.QueryRow("SELECT id, name, email FROM users WHERE name = $1", name).Scan(&user.ID, &user.Name, &user.Email)
	return
}

/**
*
* Create user
**/
func (user *User) Create() (err error) {
	statement := "insert into users (name, email, password) values ($1, $2, $3) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&user.ID)
	return
}

/**
*
* Find user by Email
**/
func (user *User) FindByEmail(email string) (err error) {

	err = database.SQL.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return
	}
	return err
}
