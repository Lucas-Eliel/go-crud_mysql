package services

import (
	"crud_mysql/internal/models"
	"crud_mysql/internal/repositories"
)

type Service interface {
	GetUsuarios() []models.User
	GetUsuario(id string) models.User
	PostUsuario(user models.User)
	DeleteUsuario(id string)
}

type ServiceImpl struct {
	repo repositories.Repository
}

func NewServiceImpl(repo repositories.Repository) Service {
	return ServiceImpl{repo}
}

func (s ServiceImpl) GetUsuarios() []models.User {
	return s.repo.FindAll()
}

func (s ServiceImpl) GetUsuario(id string) models.User {
	return s.repo.FindById(id)
}

func (s ServiceImpl) PostUsuario(user models.User) {
	s.repo.Create(user)
}

func (s ServiceImpl) DeleteUsuario(id string) {
	s.repo.Delete(id)
}
