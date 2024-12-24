package repositories

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/models"
	"gorm.io/gorm"
)

type AnnouncementRepository struct {
	DB *gorm.DB
}

func NewAnnouncementRepository(db *gorm.DB) *AnnouncementRepository {
	return &AnnouncementRepository{DB: db}
}

func (repo *AnnouncementRepository) CreateAnnouncement(announcement *models.Announcement) error {
	return repo.DB.Create(announcement).Error
}

func (repo *AnnouncementRepository) GetAnnouncementByID(id int) (*models.Announcement, error) {
	var announcement models.Announcement
	err := repo.DB.First(&announcement, id).Error
	return &announcement, err
}

func (repo *AnnouncementRepository) UpdateAnnouncement(announcement *models.Announcement) error {
	return repo.DB.Save(announcement).Error
}

func (repo *AnnouncementRepository) DeleteAnnouncement(id int) error {
	return repo.DB.Delete(&models.Announcement{}, id).Error
}

type CourseAnnouncementRepository struct {
	DB *gorm.DB
}

func NewCourseAnnouncementRepository(db *gorm.DB) *CourseAnnouncementRepository {
	return &CourseAnnouncementRepository{DB: db}
}

func (repo *CourseAnnouncementRepository) CreateCourseAnnouncement(courseAnnouncement *models.CourseAnnouncement) error {
	return repo.DB.Create(courseAnnouncement).Error
}
