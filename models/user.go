package models

import "github.com/sinclare210/Backend.git/db"

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

	result, err := stmt.Exec(user.Email,user.Password)
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