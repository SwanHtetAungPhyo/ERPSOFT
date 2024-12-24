package repositories

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (repo *CourseRepository) CreateCourse(course *models.Course) error {
	return repo.DB.Create(course).Error
}

func (repo *CourseRepository) GetCourseByID(id int) (*models.Course, error) {
	var course models.Course
	err := repo.DB.First(&course, id).Error
	return &course, err
}

func (repo *CourseRepository) UpdateCourse(course *models.Course) error {
	return repo.DB.Save(course).Error
}

func (repo *CourseRepository) DeleteCourse(id int) error {
	return repo.DB.Delete(&models.Course{}, id).Error
}
