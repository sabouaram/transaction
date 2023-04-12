package models

import (
	"errors"
	"github.com/sabouaram/transaction/sql"
	"log"
)

const (
	queryInsertUser = "INSERT INTO users(id,name,balance) VALUES(?,?,?);"
	queryGetUser    = "select * from users where id= ? ;"
	queryDeleteUser = "Delete from users where id = ? ; "
)

func (user *User) CreateUser() error {
	stmt, err := sql.Clientdb.Prepare(queryInsertUser)
	if err != nil {
		return errors.New("SQL ERROR")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name, user.Balance)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (user *User) GetUser() error {
	stmt, err := sql.Clientdb.Prepare(queryGetUser)
	if err != nil {
		return errors.New("SQL ERROR")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.Name, &user.Balance); err != nil {
		return errors.New(" SQL ERROR")
	}
	return nil
}

func (user *User) UpdateUser() error {
	_, err := sql.Clientdb.Exec("UPDATE users SET name=?, balance=?  WHERE id=?", user.Name, user.Balance, user.Id)
	if err != nil {
		return errors.New("SQL ERROR")
	}
	return nil
}

func (user *User) DeleteUser() error {
	stmt, err := sql.Clientdb.Prepare(queryDeleteUser)
	if err != nil {
		return errors.New("SQL ERROR")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Id)
	if err != nil {
		return errors.New("SQL ERROR")
		log.Println(" >> DELETE REQUEST SQL SERVICE ERROR:%v", res, err)
	}
	return nil
}
