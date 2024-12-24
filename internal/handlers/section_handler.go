package handlers

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/services"
	logging "github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SectionController struct {
	Service   *services.SectionService
	Validator *validator.Validate
	Logger    *zap.Logger
}

func NewSectionController(service *services.SectionService) *SectionController {
	return &SectionController{
		Service:   service,
		Validator: validator.New(),
		Logger:    logging.GetLogger(),
	}
}

// Register handler
func (controller *SectionController) Register(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Register handler started")
	var section models.Section
	if err := c.BodyParser(&section); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(section); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.Register(&section); err != nil {
		controller.Logger.Error("Error registering section", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Register handler completed")
	return section, nil
}

// Read handler
func (controller *SectionController) Read(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Read handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	section, err := controller.Service.GetSectionByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving section", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Read handler completed")
	return section, nil
}

// Update handler
func (controller *SectionController) Update(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Update handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	existingSection, err := controller.Service.GetSectionByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving section", zap.Error(err))
		return nil, err
	}
	var updatedSection models.Section
	if err := c.BodyParser(&updatedSection); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(updatedSection); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	existingSection.SectionName = updatedSection.SectionName
	existingSection.SectionDescription = updatedSection.SectionDescription
	existingSection.HeldBy = updatedSection.HeldBy
	if err := controller.Service.UpdateSection(existingSection); err != nil {
		controller.Logger.Error("Error updating section", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Update handler completed")
	return existingSection, nil
}

// Delete handler
func (controller *SectionController) Delete(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Delete handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.DeleteSection(id); err != nil {
		controller.Logger.Error("Error deleting section", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Delete handler completed")
	return fiber.Map{"message": "Section deleted successfully"}, nil
}

// RegisterRoutes registers the routes for the SectionController
func (controller *SectionController) RegisterRoutes(app *fiber.App) {
	sectionGroup := app.Group("/sections")
	sectionGroup.Post("/register", BaseHandle(controller.Register))
	sectionGroup.Get("/:id", BaseHandle(controller.Read))
	sectionGroup.Put("/:id", BaseHandle(controller.Update))
	sectionGroup.Delete("/:id", BaseHandle(controller.Delete))
}
