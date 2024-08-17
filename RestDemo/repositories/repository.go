package repositories

import "RestDemo/models"

type RepositoryHandler interface {
	GetStudents() ([]models.Student, error)
	UpdateStudent(student models.Student) error
	CreateStudent(student models.Student) error
	DeleteStudent(student models.Student) error
}

type Repository struct {
	RepositoryHandler RepositoryHandler
}

func NewRepository(repositoryHandler RepositoryHandler) *Repository {
	return &Repository{
		RepositoryHandler: repositoryHandler,
	}
}
