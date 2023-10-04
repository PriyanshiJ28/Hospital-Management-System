package doctor

import (
	"hospital_management_service/patient"
	"time"
)

type CreateRequest struct {
	Name      string `gorm:"column:name;type:varchar(255);not null" validator:"required" json:"name"`
	ContactNo string `gorm:"column:contact_no;type:char(10);not null" validator:"required,isContactValid" json:"contact_no"`
}
type GetRequest struct {
	ID string `gorm:"type:char(5);" validator:"required,len=5" json:"id"`
}
type UpdateRequest struct {
	ContactNo string `gorm:"column:contact_no;type:char(10);not null" validator:"required,isContactValid" json:"contact_no"`
}
type FetchPatientsByDoctorID struct {
	ID string `gorm:"column:id;type:char(5);not null" validator:"required,len=5" json:"id"`
}
type Response struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string            `json:"name"`
	ContactNo string            `json:"contact_no"`
	Patients  []patient.Patient `json:"patients"`
}
