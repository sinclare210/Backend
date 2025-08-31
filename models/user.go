package models

import (
	"errors"

	"github.com/sinclare210/Backend.git/db"
	"github.com/sinclare210/Backend.git/utils"
)

type User struct{
	Id int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (user User)Save() error{
	query := `
	INSERT INTO users(email,password)
	VALUES(?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword,err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}


	result, err := stmt.Exec(user.Email,hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = userId
	return nil
	

}

func (user *User)ValidateCredentials()error{
	query := `
	SELECT id,password FROM users WHERE email = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	rows := stmt.QueryRow(user.Email)
	var retrivedPassword string
	err = rows.Scan(&user.Id,&retrivedPassword)
	if err != nil{
		return err
	}
	
	passwordIsValid := utils.CheckPassowrHash(user.Password,retrivedPassword)

	if !passwordIsValid{
		return errors.New("credential invalid")
	}
	return nil

}