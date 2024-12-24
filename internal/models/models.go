package models

import "time"

type Role struct {
	RoleID   int    `gorm:"primaryKey;autoIncrement" json:"role_id"`
	RoleName string `gorm:"unique;not null" json:"role_name" validate:"required"`
}

type Permission struct {
	PermissionID int    `gorm:"primaryKey;autoIncrement" json:"permission_id" validate:"required"`
	Permission   string `gorm:"type:enum('CRUD','CR','CU','R');not null" json:"permission" validate:"required,oneof=CRUD CR CU R"`
}

type RolePermission struct {
	RPID         int `gorm:"primaryKey;autoIncrement" json:"rp_id" validate:"required"`
	RoleID       int `gorm:"not null" json:"role_id" validate:"required"`
	PermissionID int `gorm:"not null" json:"permission_id" validate:"required"`
}

type Employee struct {
	EmployeeID       int    `gorm:"primaryKey;autoIncrement" json:"employee_id"`
	EmployeeName     string `gorm:"not null" json:"employee_name" validate:"required"`
	EmployeePassword string `gorm:"not null" json:"employee_password" validate:"required"`
	EmployeeEmail    string `gorm:"unique;not null" json:"employee_email" validate:"required,email"`
	RoleID           int    `gorm:"not null" json:"role_id" validate:"required"`
}

type Student struct {
	StudentID   int    `gorm:"primaryKey;autoIncrement" json:"student_id"`
	StudentName string `gorm:"not null" json:"student_name" validate:"required"`
	Email       string `gorm:"unique;not null" json:"email" validate:"required,email"`
}

type Course struct {
	CourseID      int       `gorm:"primaryKey;autoIncrement" json:"course_id"`
	CourseName    string    `gorm:"not null" json:"course_name" validate:"required"`
	CreatedBy     int       `gorm:"not null" json:"created_by" validate:"required"`
	Description   string    `gorm:"not null" json:"description" validate:"required"`
	StartDate     string    `gorm:"type:date" json:"start_date"`
	EndDate       string    `gorm:"type:date" json:"end_date"`
	LearnPlatform string    `gorm:"type:varchar(150)" json:"learn_platform"`
	CreateAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"create_at"`
}

type Announcement struct {
	AnnouncementID          int    `gorm:"primaryKey;autoIncrement" json:"announcement_id"`
	AnnouncementDescription string `gorm:"not null" json:"announcement_description" validate:"required"`
	CreatedBy               int    `gorm:"not null" json:"created_by" validate:"required"`
}

type CourseAnnouncement struct {
	CourseAnnouncementID int `gorm:"primaryKey;autoIncrement" json:"course_announcement_id"`
	CourseID             int `gorm:"not null" json:"course_id" validate:"required"`
	AnnouncementID       int `gorm:"not null" json:"announcement_id" validate:"required"`
}

type CourseStudent struct {
	CourseStudentID int `gorm:"primaryKey;autoIncrement" json:"course_student_id"`
	CourseID        int `gorm:"not null" json:"course_id" validate:"required"`
	StudentID       int `gorm:"not null" json:"student_id" validate:"required"`
}

type Section struct {
	SectionID          int    `gorm:"primaryKey;autoIncrement" json:"section_id"`
	SectionName        string `gorm:"not null" json:"section_name" validate:"required"`
	SectionDescription string `gorm:"not null" json:"section_description" validate:"required"`
	HeldBy             int    `gorm:"not null" json:"held_by" validate:"required"`
}

type SectionBooking struct {
	BookingID   int    `gorm:"primaryKey;autoIncrement" json:"booking_id"`
	SectionID   int    `gorm:"not null" json:"section_id" validate:"required"`
	UserID      int    `gorm:"not null" json:"user_id" validate:"required"`
	BookingDate string `gorm:"type:datetime;not null" json:"booking_date" validate:"required"`
}
