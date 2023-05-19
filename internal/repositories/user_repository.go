package repositories

import (
	"context"
	"crud_mysql/internal/models"
	"database/sql"
)

type Repository interface {
	FindAll() []models.User
	FindById(id string) models.User
	Create(user models.User)
	Delete(id string)
}

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(db *sql.DB) Repository {
	return RepositoryImpl{db}
}

func (r RepositoryImpl) FindAll() []models.User {

	results, err := r.db.Query("SELECT * FROM user")

	if err != nil {
		panic(err)
	}

	var users []models.User

	for results.Next() {
		var user models.User
		erro := results.Scan(&user.ID, &user.Nome)

		if erro != nil {
			panic(erro)
		}
		users = append(users, user)
	}

	return users
}

func (r RepositoryImpl) FindById(id string) models.User {
	results, err := r.db.Query("SELECT * FROM user WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	var users []models.User

	for results.Next() {
		var user models.User
		erro := results.Scan(&user.ID, &user.Nome)

		if erro != nil {
			panic(erro)
		}
		users = append(users, user)
	}

	return users[0]
}

func (r RepositoryImpl) Create(user models.User) {
	_, err := r.db.ExecContext(context.Background(), "INSERT INTO user(ID, NOME) VALUES(?, ?)", user.ID, user.Nome)

	if err != nil {
		panic(err)
	}
}

func (r RepositoryImpl) Delete(id string) {
	_, err := r.db.ExecContext(context.Background(), "DELETE FROM user WHERE id = ?", id)

	if err != nil {
		panic(err)
	}
}
