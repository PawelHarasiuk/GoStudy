package repositories

import (
	"RestDemo/models"
)

type RepositoryHandler interface {
	Load() map[string]models.Student
	Save(students map[string]models.Student)
}

type Repository struct {
	RepositoryHandler RepositoryHandler
}

func NewRepository(repositoryHandler RepositoryHandler) *Repository {
	return &Repository{
		RepositoryHandler: repositoryHandler,
	}
}
