package services

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/repo"
)

// RoleService struct
type RoleService struct {
	Repo *repositories.RoleRepository
}

// NewRoleService creates a new RoleService
func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

// CreateRole adds a new role to the database
func (service *RoleService) CreateRole(role *models.Role) error {
	return service.Repo.CreateRole(role)
}

// GetRoles retrieves all roles from the database
func (service *RoleService) GetRoles(roles *[]models.Role) error {
	return service.Repo.GetRoles(roles)
}

// GetRoleByID retrieves a role by its ID from the database
func (service *RoleService) GetRoleByID(role *models.Role, id int) error {
	return service.Repo.GetRoleByID(role, id)
}

// UpdateRole updates an existing role in the database
func (service *RoleService) UpdateRole(role *models.Role) error {
	return service.Repo.UpdateRole(role)
}

// DeleteRole removes a role from the database
func (service *RoleService) DeleteRole(id int) error {
	return service.Repo.DeleteRole(id)
}
