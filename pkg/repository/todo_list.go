package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id)", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
	}

	createUserLisQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id)", userListTable)
	_, err = tx.Exec(createUserLisQuery, userId, id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
	}

	return id, tx.Commit()
}
