package handlers

import (
	"time"

	middlewares "github.com/SwanHtetAungPhyo/esfor/internal/middleware"
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/services"
	logging "github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type EmployeeController struct {
	Service   *services.EmployeeService
	Validator *validator.Validate
	Logger    *zap.Logger
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		Service:   service,
		Validator: validator.New(),
		Logger:    logging.GetLogger(),
	}
}

// Register handler
func (controller *EmployeeController) Register(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Register handler started")
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(employee); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.Register(&employee); err != nil {
		controller.Logger.Error("Error registering employee", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Register handler completed")
	return employee, nil
}

// Login handler
func (controller *EmployeeController) Login(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Login handler started")
	var loginData struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(loginData); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	token, err := controller.Service.Login(loginData.Email, loginData.Password)
	if err != nil {
		controller.Logger.Error("Error logging in", zap.Error(err))
		return nil, err
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Domain:   "localhost",
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "None",
	})
	controller.Logger.Info("Login handler completed")
	return fiber.Map{"token": token}, nil
}

// Read handler
func (controller *EmployeeController) Read(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Read handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	employee, err := controller.Service.GetEmployeeByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving employee", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Read handler completed")
	return employee, nil
}

// Delete handler
func (controller *EmployeeController) Delete(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Delete handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.DeleteEmployee(id); err != nil {
		controller.Logger.Error("Error deleting employee", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Delete handler completed")
	return fiber.Map{"message": "Employee deleted successfully"}, nil
}

// RegisterRoutes registers the routes for the EmployeeController
func (controller *EmployeeController) RegisterRoutes(app *fiber.App) {
	app.Post("/employees/register", BaseHandle(controller.Register))
	app.Post("/employees/login", middlewares.LoggingMiddleware(), BaseHandle(controller.Login))
	app.Get("/employees/:id", middlewares.JwtMiddleware(), middlewares.LoggingMiddleware(), BaseHandle(controller.Read))
	app.Delete("/employees/:id", middlewares.JwtMiddleware(), BaseHandle(controller.Delete))
}
