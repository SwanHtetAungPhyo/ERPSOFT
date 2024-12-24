package repositories

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (repo *StudentRepository) CreateStudent(student *models.Student) error {
	return repo.DB.Create(student).Error
}

func (repo *StudentRepository) GetStudentByEmail(email string) (*models.Student, error) {
	var student models.Student
	err := repo.DB.Where("email = ?", email).First(&student).Error
	return &student, err
}

func (repo *StudentRepository) GetStudentByID(id int) (*models.Student, error) {
	var student models.Student
	err := repo.DB.First(&student, id).Error
	return &student, err
}

func (repo *StudentRepository) UpdateStudent(student *models.Student) error {
	return repo.DB.Save(student).Error
}

func (repo *StudentRepository) DeleteStudent(id int) error {
	return repo.DB.Delete(&models.Student{}, id).Error
}
