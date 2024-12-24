package services

import (
	"errors"

	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	repositories "github.com/SwanHtetAungPhyo/esfor/internal/repo"
	"gorm.io/gorm"
)

type AnnouncementService struct {
	Repo                   *repositories.AnnouncementRepository
	CourseAnnouncementRepo *repositories.CourseAnnouncementRepository
	EmployeeRepo           *repositories.EmployeeRepository
	DB                     *gorm.DB
}

func NewAnnouncementService(repo *repositories.AnnouncementRepository, courseAnnouncementRepo *repositories.CourseAnnouncementRepository, employeeRepo *repositories.EmployeeRepository, db *gorm.DB) *AnnouncementService {
	return &AnnouncementService{Repo: repo, CourseAnnouncementRepo: courseAnnouncementRepo, EmployeeRepo: employeeRepo, DB: db}
}

func (service *AnnouncementService) Register(announcement *models.Announcement, courseID int) error {
	return service.DB.Transaction(func(tx *gorm.DB) error {
		if _, err := service.EmployeeRepo.GetEmployeeByID(announcement.CreatedBy); err != nil {
			return errors.New("invalid created_by: employee not found")
		}
		if err := service.Repo.CreateAnnouncement(announcement); err != nil {
			return err
		}
		courseAnnouncement := &models.CourseAnnouncement{
			CourseID:       courseID,
			AnnouncementID: announcement.AnnouncementID,
		}
		if err := service.CourseAnnouncementRepo.CreateCourseAnnouncement(courseAnnouncement); err != nil {
			return err
		}
		return nil
	})
}

func (service *AnnouncementService) GetAnnouncementByID(id int) (*models.Announcement, error) {
	return service.Repo.GetAnnouncementByID(id)
}

func (service *AnnouncementService) UpdateAnnouncement(announcement *models.Announcement) error {
	if _, err := service.EmployeeRepo.GetEmployeeByID(announcement.CreatedBy); err != nil {
		return errors.New("invalid created_by: employee not found")
	}
	return service.Repo.UpdateAnnouncement(announcement)
}

func (service *AnnouncementService) DeleteAnnouncement(id int) error {
	return service.Repo.DeleteAnnouncement(id)
}
