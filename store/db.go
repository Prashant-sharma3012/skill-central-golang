package store

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	DB       *sql.DB
	Employee *EmployeeStore
}

var DB *sql.DB

func Connect(driver string, connectionUrl string) *Repository {
	DB, err := sql.Open()

	if err != nil {
		fmt.Println("SQL Connection Error", err)
	}

	return &Repository{
		DB:       DB,
		Employee: &EmployeeStore{},
	}
}
