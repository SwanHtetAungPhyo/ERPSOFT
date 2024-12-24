package handlers

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/services"
	logging "github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CourseController struct {
	Service   *services.CourseService
	Validator *validator.Validate
	Logger    *zap.Logger
}

func NewCourseController(service *services.CourseService) *CourseController {
	return &CourseController{
		Service:   service,
		Validator: validator.New(),
		Logger:    logging.GetLogger(),
	}
}

// Register handler
func (controller *CourseController) Register(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Register handler started")
	var course models.Course
	if err := c.BodyParser(&course); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(course); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.Register(&course); err != nil {
		controller.Logger.Error("Error registering course", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Register handler completed")
	return course, nil
}

// Read handler
func (controller *CourseController) Read(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Read handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	course, err := controller.Service.GetCourseByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving course", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Read handler completed")
	return course, nil
}

// Update handler
func (controller *CourseController) Update(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Update handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	existingCourse, err := controller.Service.GetCourseByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving course", zap.Error(err))
		return nil, err
	}
	var updatedCourse models.Course
	if err := c.BodyParser(&updatedCourse); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(updatedCourse); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	existingCourse.CourseName = updatedCourse.CourseName
	existingCourse.CreatedBy = updatedCourse.CreatedBy
	existingCourse.Description = updatedCourse.Description
	existingCourse.StartDate = updatedCourse.StartDate
	existingCourse.EndDate = updatedCourse.EndDate
	existingCourse.LearnPlatform = updatedCourse.LearnPlatform
	if err := controller.Service.UpdateCourse(existingCourse); err != nil {
		controller.Logger.Error("Error updating course", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Update handler completed")
	return existingCourse, nil
}

// Delete handler
func (controller *CourseController) Delete(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Delete handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.DeleteCourse(id); err != nil {
		controller.Logger.Error("Error deleting course", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Delete handler completed")
	return fiber.Map{"message": "Course deleted successfully"}, nil
}

// RegisterRoutes registers the routes for the CourseController
func (controller *CourseController) RegisterRoutes(app *fiber.App) {
	courseGroup := app.Group("/courses")
	courseGroup.Post("/register", BaseHandle(controller.Register))
	courseGroup.Get("/:id", BaseHandle(controller.Read))
	courseGroup.Put("/:id", BaseHandle(controller.Update))
	courseGroup.Delete("/:id", BaseHandle(controller.Delete))
}
