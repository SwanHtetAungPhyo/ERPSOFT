package handlers

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/services"
	logging "github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AnnouncementController struct {
	Service   *services.AnnouncementService
	Validator *validator.Validate
	Logger    *zap.Logger
}

func NewAnnouncementController(service *services.AnnouncementService) *AnnouncementController {
	return &AnnouncementController{
		Service:   service,
		Validator: validator.New(),
		Logger:    logging.GetLogger(),
	}
}

// Register godoc
// @Summary Register a new announcement
// @Description Register a new announcement
// @Tags announcements
// @Accept json
// @Produce json
// @Param announcement body models.Announcement true "Announcement"
// @Param course_id body int true "Course ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /announcements/register [post]
func (controller *AnnouncementController) Register(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Register handler started")
	var request struct {
		Announcement models.Announcement `json:"announcement"`
		CourseID     int                 `json:"course_id" validate:"required"`
	}
	if err := c.BodyParser(&request); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(request); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.Register(&request.Announcement, request.CourseID); err != nil {
		controller.Logger.Error("Error registering announcement", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Register handler completed")
	return request.Announcement, nil
}

// Read godoc
// @Summary Get an announcement by ID
// @Description Get an announcement by ID
// @Tags announcements
// @Accept json
// @Produce json
// @Param id path int true "Announcement ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /announcements/{id} [get]
func (controller *AnnouncementController) Read(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Read handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	announcement, err := controller.Service.GetAnnouncementByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving announcement", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Read handler completed")
	return announcement, nil
}

// Update godoc
// @Summary Update an announcement
// @Description Update an announcement
// @Tags announcements
// @Accept json
// @Produce json
// @Param id path int true "Announcement ID"
// @Param announcement body models.Announcement true "Announcement"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /announcements/{id} [put]
func (controller *AnnouncementController) Update(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Update handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	existingAnnouncement, err := controller.Service.GetAnnouncementByID(id)
	if err != nil {
		controller.Logger.Error("Error retrieving announcement", zap.Error(err))
		return nil, err
	}
	var updatedAnnouncement models.Announcement
	if err := c.BodyParser(&updatedAnnouncement); err != nil {
		controller.Logger.Error("Error parsing body", zap.Error(err))
		return nil, err
	}
	if err := controller.Validator.Struct(updatedAnnouncement); err != nil {
		controller.Logger.Error("Validation error", zap.Error(err))
		return nil, err
	}
	existingAnnouncement.AnnouncementDescription = updatedAnnouncement.AnnouncementDescription
	existingAnnouncement.CreatedBy = updatedAnnouncement.CreatedBy
	if err := controller.Service.UpdateAnnouncement(existingAnnouncement); err != nil {
		controller.Logger.Error("Error updating announcement", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Update handler completed")
	return existingAnnouncement, nil
}

// Delete godoc
// @Summary Delete an announcement
// @Description Delete an announcement
// @Tags announcements
// @Accept json
// @Produce json
// @Param id path int true "Announcement ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /announcements/{id} [delete]
func (controller *AnnouncementController) Delete(c *fiber.Ctx) (interface{}, error) {
	controller.Logger.Info("Delete handler started")
	id, err := c.ParamsInt("id")
	if err != nil {
		controller.Logger.Error("Error parsing ID", zap.Error(err))
		return nil, err
	}
	if err := controller.Service.DeleteAnnouncement(id); err != nil {
		controller.Logger.Error("Error deleting announcement", zap.Error(err))
		return nil, err
	}
	controller.Logger.Info("Delete handler completed")
	return fiber.Map{"message": "Announcement deleted successfully"}, nil
}

// RegisterRoutes registers the routes for the AnnouncementController
func (controller *AnnouncementController) RegisterRoutes(app *fiber.App) {
	announcementGroup := app.Group("/announcements")
	announcementGroup.Post("/register", BaseHandle(controller.Register))
	announcementGroup.Get("/:id", BaseHandle(controller.Read))
	announcementGroup.Put("/:id", BaseHandle(controller.Update))
	announcementGroup.Delete("/:id", BaseHandle(controller.Delete))
}
