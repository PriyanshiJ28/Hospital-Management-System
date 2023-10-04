package doctor

import (
	"gorm.io/gorm"
	"hospital_management_service/constants"
	"hospital_management_service/patient"
)

type IRepository interface {
	Create(doctor *Doctor)
	Get(id uint) (*Doctor, error)
	Update(doctor *Doctor)
	FetchPatientsByDoctorID(doctorID string) []patient.Response
}
type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
func (r *Repository) Create(doctor *Doctor) error {
	return r.db.Create(doctor).Error
}
func (r *Repository) Get(id string) (*Doctor, error) {
	var doctor Doctor
	if err := r.db.Where(constants.ID+" = ? ", id).Preload("Patients").Find(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}
func (r *Repository) Update(doctor *Doctor) (*Doctor, error) {
	return doctor, r.db.Save(doctor).Error
}

func (r *Repository) FetchPatientsByDoctorID(doctorID string) ([]patient.Patient, error) {
	var doctor Doctor
	if err := r.db.Where(constants.ID+" = ? ", doctorID).Preload("Patients").Find(&doctor).Error; err != nil {
		return nil, err
	}
	return doctor.Patients, nil
}
