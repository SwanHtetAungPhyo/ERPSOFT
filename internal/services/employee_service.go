package services

import (
	"errors"
	"time"

	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	repositories "github.com/SwanHtetAungPhyo/esfor/internal/repo"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeService struct {
	Repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (service *EmployeeService) Register(employee *models.Employee) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.EmployeePassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	employee.EmployeePassword = string(hashedPassword)
	return service.Repo.CreateEmployee(employee)
}

func (service *EmployeeService) Login(email, password string) (string, error) {
	employee, err := service.Repo.GetEmployeeByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(employee.EmployeePassword), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}
	token, err := generateJWT(employee)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *EmployeeService) GetEmployeeByID(id int) (*models.Employee, error) {
	return service.Repo.GetEmployeeByID(id)
}

func (service *EmployeeService) DeleteEmployee(id int) error {
	return service.Repo.DeleteEmployee(id)
}

func generateJWT(employee *models.Employee) (string, error) {
	claims := jwt.MapClaims{
		"employee_id": employee.EmployeeID,
		"email":       employee.EmployeeEmail,
		"exp":         time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("p0EKBU4NvrdYGqnYgCNzvdZQHNrZiUjj4jQ7ZVrgRj5Ymg5S")) // Replace with your actual secret key
}
