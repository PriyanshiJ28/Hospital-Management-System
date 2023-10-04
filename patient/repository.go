package patient

import (
	"gorm.io/gorm"
	"hospital_management_service/constants"
)

type IRepository interface {
	Create(doctor *Patient)
	Get(id string) (*Patient, error)
	Update(doctor *Patient) error
	FetchPatientsByDoctorID(doctorID string) ([]*Patient, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
func (r *Repository) Create(patient *Patient) error {
	return r.db.Create(patient).Error
}
func (r *Repository) Get(id string) (*Patient, error) {
	var patient Patient
	if err := r.db.Where(constants.ID+" = ? ", id).First(&patient).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}
func (r *Repository) Update(patient *Patient) error {
	return r.db.Save(patient).Error
}
