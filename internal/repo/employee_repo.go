package repositories

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (repo *EmployeeRepository) CreateEmployee(employee *models.Employee) error {
	return repo.DB.Create(employee).Error
}

func (repo *EmployeeRepository) GetEmployeeByEmail(email string) (*models.Employee, error) {
	var employee models.Employee
	err := repo.DB.Where("employee_email = ?", email).First(&employee).Error
	return &employee, err
}

func (repo *EmployeeRepository) GetEmployeeByID(id int) (*models.Employee, error) {
	var employee models.Employee
	err := repo.DB.First(&employee, id).Error
	return &employee, err
}

func (repo *EmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	return repo.DB.Save(employee).Error
}

func (repo *EmployeeRepository) DeleteEmployee(id int) error {
	return repo.DB.Delete(&models.Employee{}, id).Error
}
