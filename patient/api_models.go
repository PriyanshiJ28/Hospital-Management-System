package patient

import "time"

type CreateRequest struct {
	Name      string `gorm:"column:name;type:varchar(255);not null" validate:"required" json:"name"`
	ContactNo string `gorm:"column:contact_no;type:char(10);not null"  validate:"required,isContactValid" json:"contact_no"`
	Address   string `gorm:"column:address;type:varchar(255);not null" validate:"required" json:"address"`
	DoctorID  string `gorm:"column:doctor_id;type:char(5);not null;" validate:"required,len=5" json:"doctor_id"`
}
type GetRequest struct {
	ID string `gorm:"type:char(5);primary_key;" validate:"required,len=5"`
}
type UpdateRequest struct {
	ContactNo string `gorm:"column:contact_no;type:char(10);not null" validate:"required,isContactValid" json:"contact_no"`
	Address   string `gorm:"column:address;type:varchar(255);not null" validate:"required" json:"address"`
	DoctorID  string `gorm:"column:doctor_id;type:char(5);not null" json:"doctor_id"`
}
type Response struct {
	ID        string `gorm:"type:char(5);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	ContactNo string `gorm:"column:contact_no;type:char(10);not null" json:"contact_no"`
	Address   string `gorm:"column:address;type:varchar(255);not null" json:"address"`
	DoctorID  string `gorm:"column:doctor_id;type:char(5);not null;foreignKey" json:"doctor_id"`
}
