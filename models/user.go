package models

import (
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