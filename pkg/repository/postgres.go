package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable      = "users"
	todoListsTable = "todo_lists"
	userListsTable = "users_lists"
	todoItemsTable = "todo_items"
	ListsItemTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresBD(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password, conf.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
