package patient

import (
	"time"
)

type Patient struct {
	ID        string `gorm:"type:char(5);primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	ContactNo string `gorm:"column:contact_no;type:char(10);not null" json:"contact_no"`
	Address   string `gorm:"column:address;type:varchar(255);not null" json:"address"`
	DoctorID  string `gorm:"column:doctor_id;type:char(5);not null;" json:"doctor_id"`
}
