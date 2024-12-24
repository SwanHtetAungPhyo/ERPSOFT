package services

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/repo"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{Repo: repo}
}

func (service *StudentService) Register(student *models.Student) error {
	return service.Repo.CreateStudent(student)
}

func (service *StudentService) GetStudentByID(id int) (*models.Student, error) {
	return service.Repo.GetStudentByID(id)
}

func (service *StudentService) UpdateStudent(student *models.Student) error {
	return service.Repo.UpdateStudent(student)
}

func (service *StudentService) DeleteStudent(id int) error {
	return service.Repo.DeleteStudent(id)
}
