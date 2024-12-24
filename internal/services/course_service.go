package services

import (
	"errors"

	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	repositories "github.com/SwanHtetAungPhyo/esfor/internal/repo"
)

type CourseService struct {
	Repo         *repositories.CourseRepository
	EmployeeRepo *repositories.EmployeeRepository
}

func NewCourseService(repo *repositories.CourseRepository, employeeRepo *repositories.EmployeeRepository) *CourseService {
	return &CourseService{Repo: repo, EmployeeRepo: employeeRepo}
}

func (service *CourseService) Register(course *models.Course) error {
	if _, err := service.EmployeeRepo.GetEmployeeByID(course.CreatedBy); err != nil {
		return errors.New("invalid created_by: employee not found")
	}
	return service.Repo.CreateCourse(course)
}

func (service *CourseService) GetCourseByID(id int) (*models.Course, error) {
	return service.Repo.GetCourseByID(id)
}

func (service *CourseService) UpdateCourse(course *models.Course) error {
	if _, err := service.EmployeeRepo.GetEmployeeByID(course.CreatedBy); err != nil {
		return errors.New("invalid created_by: employee not found")
	}
	return service.Repo.UpdateCourse(course)
}

func (service *CourseService) DeleteCourse(id int) error {
	return service.Repo.DeleteCourse(id)
}
