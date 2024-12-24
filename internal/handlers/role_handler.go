package handlers

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	Service   *services.RoleService
	Validator *validator.Validate
}

func NewRoleController(service *services.RoleService) *RoleController {
	return &RoleController{
		Service:   service,
		Validator: validator.New(),
	}
}

// CreateRole handler
func (controller *RoleController) CreateRole(c *fiber.Ctx) (interface{}, error) {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return nil, err
	}
	if err := controller.Validator.Struct(role); err != nil {
		return nil, err
	}
	if err := controller.Service.CreateRole(&role); err != nil {
		return nil, err
	}
	return role, nil
}

// GetRoles handler
func (controller *RoleController) GetRoles(c *fiber.Ctx) (interface{}, error) {
	var roles []models.Role
	if err := controller.Service.GetRoles(&roles); err != nil {
		return nil, err
	}
	return roles, nil
}

// GetRoleByID handler
func (controller *RoleController) GetRoleByID(c *fiber.Ctx) (interface{}, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return nil, err
	}
	var role models.Role
	if err := controller.Service.GetRoleByID(&role, id); err != nil {
		return nil, err
	}
	return role, nil
}

// UpdateRole handler
func (controller *RoleController) UpdateRole(c *fiber.Ctx) (interface{}, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return nil, err
	}
	var role models.Role
	if err := controller.Service.GetRoleByID(&role, id); err != nil {
		return nil, err
	}
	if err := c.BodyParser(&role); err != nil {
		return nil, err
	}
	if err := controller.Validator.Struct(role); err != nil {
		return nil, err
	}
	if err := controller.Service.UpdateRole(&role); err != nil {
		return nil, err
	}
	return role, nil
}

// DeleteRole handler
func (controller *RoleController) DeleteRole(c *fiber.Ctx) (interface{}, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return nil, err
	}
	if err := controller.Service.DeleteRole(id); err != nil {
		return nil, err
	}
	return fiber.Map{"message": "Role deleted successfully"}, nil
}

// RegisterRoutes registers the routes for the RoleController
func (controller *RoleController) RegisterRoutes(app *fiber.App) {
	app.Post("/roles", BaseHandle(controller.CreateRole))
	app.Get("/roles", BaseHandle(controller.GetRoles))
	app.Get("/roles/:id", BaseHandle(controller.GetRoleByID))
	app.Put("/roles/:id", BaseHandle(controller.UpdateRole))
	app.Delete("/roles/:id", BaseHandle(controller.DeleteRole))
}
