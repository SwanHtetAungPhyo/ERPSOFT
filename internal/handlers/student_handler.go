package handlers

import (
	middlewares "github.com/SwanHtetAungPhyo/esfor/internal/middleware"
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/services"
	logging "github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StudentController struct {
	Service   *services.StudentService
	Validator *validator.Validate
	Logger    *zap.Logger
}

func NewStudentController(service *services.StudentService) *StudentController {
	return &StudentController{
		Service:   service,
		Validator: validator.New(),
		Logger:    logging.GetLogger(),
	}
}

// Register handler
func (controller *StudentController) Register(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Register handler started")
	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(student); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.Register(&student); err != nil {
		controller.Logger.Error("Error registering student", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Register handler completed")
	return student, nil
}

// Read handler
func (controller *StudentController) Read(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Read handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	student, err := controller.Service.GetStudentByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving student", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Read handler completed")
	return student, nil
}

// Update handler
func (controller *StudentController) Update(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Update handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	existingStudent, err := controller.Service.GetStudentByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving student", zap.Error(err))
		return nil, err
	}
	var updatedStudent models.Student
	if err := c.BodyParser(&updatedStudent); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(updatedStudent); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	existingStudent.StudentName = updatedStudent.StudentName
	existingStudent.Email = updatedStudent.Email
	if err := controller.Service.UpdateStudent(existingStudent); err != nil {
		controller.Logger.Error("Error updating student", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Update handler completed")
	return existingStudent, nil
}

// Delete handler
func (controller *StudentController) Delete(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Delete handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.DeleteStudent(id); err != nil {
		controller.Logger.Error("Error deleting student", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Delete handler completed")
	return fiber.Map{"message": "Student deleted successfully"}, nil
}

// RegisterRoutes registers the routes for the StudentController
func (controller *StudentController) RegisterRoutes(app *fiber.App) {
	studentGroup := app.Group("/students")
	studentGroup.Use(middlewares.JwtMiddleware())
	studentGroup.Post("/register", BaseHandle(controller.Register))
	app.Get("/students/:id", BaseHandle(controller.Read))
	app.Put("/students/:id", BaseHandle(controller.Update))
	app.Delete("/students/:id", BaseHandle(controller.Delete))
}
