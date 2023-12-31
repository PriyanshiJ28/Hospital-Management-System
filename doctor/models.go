package doctor

import (
	"hospital_management_service/patient"
	"time"
)

type Doctor struct {
	ID        string `gorm:"type:char(5);primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string            `gorm:"column:name;type:varchar(255);not null" json:"name"`
	ContactNo string            `gorm:"column:contact_no;type:char(10);not null" json:"contact_no"`
	Patients  []patient.Patient `gorm:"one2many:doctor_id" json:"patients"`
}
