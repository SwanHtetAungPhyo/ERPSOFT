package repositories

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// RoleRepository struct
type RoleRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

// NewRoleRepository creates a new RoleRepository
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		DB:     db,
		Logger: utils.GetLogger(),
	}
}

// CreateRole adds a new role to the database
func (repo *RoleRepository) CreateRole(role *models.Role) error {
	repo.Logger.Info("Creating new role", zap.String("role_name", role.RoleName))
	if err := repo.DB.Create(role).Error; err != nil {
		repo.Logger.Error("Failed to create role", zap.Error(err))
		return err
	}
	repo.Logger.Info("Role created successfully", zap.Int("role_id", role.RoleID))
	return nil
}

// GetRoles retrieves all roles from the database
func (repo *RoleRepository) GetRoles(roles *[]models.Role) error {
	repo.Logger.Info("Retrieving all roles")
	if err := repo.DB.Find(roles).Error; err != nil {
		repo.Logger.Error("Failed to retrieve roles", zap.Error(err))
		return err
	}
	repo.Logger.Info("Roles retrieved successfully", zap.Int("count", len(*roles)))
	return nil
}

// GetRoleByID retrieves a role by its ID from the database
func (repo *RoleRepository) GetRoleByID(role *models.Role, id int) error {
	repo.Logger.Info("Retrieving role by ID", zap.Int("role_id", id))
	if err := repo.DB.First(role, id).Error; err != nil {
		repo.Logger.Error("Failed to retrieve role", zap.Error(err))
		return err
	}
	repo.Logger.Info("Role retrieved successfully", zap.Int("role_id", role.RoleID))
	return nil
}

// UpdateRole updates an existing role in the database
func (repo *RoleRepository) UpdateRole(role *models.Role) error {
	repo.Logger.Info("Updating role", zap.Int("role_id", role.RoleID))
	if err := repo.DB.Save(role).Error; err != nil {
		repo.Logger.Error("Failed to update role", zap.Error(err))
		return err
	}
	repo.Logger.Info("Role updated successfully", zap.Int("role_id", role.RoleID))
	return nil
}

// DeleteRole removes a role from the database
func (repo *RoleRepository) DeleteRole(id int) error {
	repo.Logger.Info("Deleting role", zap.Int("role_id", id))
	if err := repo.DB.Delete(&models.Role{}, id).Error; err != nil {
		repo.Logger.Error("Failed to delete role", zap.Error(err))
		return err
	}
	repo.Logger.Info("Role deleted successfully", zap.Int("role_id", id))
	return nil
}
