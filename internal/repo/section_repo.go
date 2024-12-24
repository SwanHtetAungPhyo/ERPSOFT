package repositories

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"gorm.io/gorm"
)

type SectionRepository struct {
	DB *gorm.DB
}

func NewSectionRepository(db *gorm.DB) *SectionRepository {
	return &SectionRepository{DB: db}
}

func (repo *SectionRepository) CreateSection(section *models.Section) error {
	return repo.DB.Create(section).Error
}

func (repo *SectionRepository) GetSectionByID(id int) (*models.Section, error) {
	var section models.Section
	err := repo.DB.First(&section, id).Error
	return &section, err
}

func (repo *SectionRepository) UpdateSection(section *models.Section) error {
	return repo.DB.Save(section).Error
}

func (repo *SectionRepository) DeleteSection(id int) error {
	return repo.DB.Delete(&models.Section{}, id).Error
}
