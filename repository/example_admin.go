package repository

import (
	"database/sql"
	"log"

	"github.com/Uchel/auth-jwt2/model"
)

type ExampleAdminLoginRepo interface {
	GetByEmailExampleAdmin(email string) (string, string)
}

type exampleAdminLoginRepo struct {
	db *sql.DB
}

func (u *exampleAdminLoginRepo) GetByEmailExampleAdmin(email string) (string, string) {
	var exampleAdmin model.ExampleAdmin

	query := "SELECT email, password FROM example_admin WHERE email=$1"

	row := u.db.QueryRow(query, email)

	if err := row.Scan(&exampleAdmin.Email, &exampleAdmin.Password); err != nil {
		log.Println(err)
	}

	if exampleAdmin.Email == "" {
		return "user not found", " or password uncorrect"
	}

	return exampleAdmin.Email, exampleAdmin.Password
}
func NewExampleAdminLoginRepo(db *sql.DB) ExampleAdminLoginRepo {
	repo := new(exampleAdminLoginRepo)

	repo.db = db

	return repo
}
