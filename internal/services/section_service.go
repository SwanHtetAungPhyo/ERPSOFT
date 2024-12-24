package services

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	repositories "github.com/SwanHtetAungPhyo/esfor/internal/repo"
)

type SectionService struct {
	Repo *repositories.SectionRepository
}

func NewSectionService(repo *repositories.SectionRepository) *SectionService {
	return &SectionService{Repo: repo}
}

func (service *SectionService) Register(section *models.Section) error {
	return service.Repo.CreateSection(section)
}

func (service *SectionService) GetSectionByID(id int) (*models.Section, error) {
	return service.Repo.GetSectionByID(id)
}

func (service *SectionService) UpdateSection(section *models.Section) error {
	return service.Repo.UpdateSection(section)
}

func (service *SectionService) DeleteSection(id int) error {
	return service.Repo.DeleteSection(id)
}
