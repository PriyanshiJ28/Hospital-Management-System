package doctor

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gopkg.in/go-playground/validator.v9"
	"hospital_management_service/common.go"
	"hospital_management_service/patient"
)

type IService interface {
	Create(request CreateRequest)
	Get(request GetRequest) Response
	Update(request UpdateRequest)
	FetchPatientsByDoctorID(request FetchPatientsByDoctorID) []patient.Response
}

type Service struct {
	repo      Repository
	validator *validator.Validate
}

func NewService(repo Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:      repo,
		validator: validate,
	}
}
func (s *Service) Create(request CreateRequest) error {
	err := s.validator.RegisterValidation("isContactValid", common_go.ValidateContactNo)
	if err != nil {
		fmt.Println("Error in validation:", err.Error())
	}
	err = s.validator.Struct(request)
	if err != nil {
		panic("Invalid Entries")
	}
	id, err := gonanoid.Generate(`"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"`, 5)
	doctorModel := Doctor{
		ID:        id,
		Name:      request.Name,
		ContactNo: request.ContactNo,
	}
	err = s.repo.Create(&doctorModel)
	return err
}
func (s *Service) Get(request GetRequest) (Response, error) {
	err := s.validator.Struct(request)
	if err != nil {
		panic("Invalid ID")
	}
	doc, err2 := s.repo.Get(request.ID)
	response := Response{
		ID:        doc.ID,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
		Name:      doc.Name,
		ContactNo: doc.ContactNo,
		Patients:  doc.Patients,
	}
	return response, err2
}
func (s *Service) Update(id string, request UpdateRequest) error {
	err := s.validator.RegisterValidation("isContactValid", common_go.ValidateContactNo)
	if err != nil {
		fmt.Println("Error in validation:", err.Error())
	}
	err = s.validator.Struct(request)
	if err != nil {
		panic("Invalid Entries")
	}
	doc, err1 := s.repo.Get(id)
	if err1 != nil {
		panic("Cant get repo's doc")
	}
	doc.ContactNo = request.ContactNo
	doc, err = s.repo.Update(doc)
	return err
}
func (s *Service) FetchPatientsByDoctorID(request FetchPatientsByDoctorID) ([]patient.Response, error) {
	err := s.validator.Struct(request)
	if err != nil {
		panic("Invalid Entries")
	}
	patients, err := s.repo.FetchPatientsByDoctorID(request.ID)
	if err != nil {
		return nil, err
	}
	var response []patient.Response
	for _, pat := range patients {
		p := patient.Response{
			ID:        pat.ID,
			CreatedAt: pat.CreatedAt,
			UpdatedAt: pat.UpdatedAt,
			Name:      pat.Name,
			ContactNo: pat.ContactNo,
			Address:   pat.Address,
			DoctorID:  pat.DoctorID,
		}
		response = append(response, p)
	}
	return response, nil
}
