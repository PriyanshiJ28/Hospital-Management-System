package patient

import (
	"fmt"
	"hospital_management_service/common.go"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gopkg.in/go-playground/validator.v9"
)

type IService interface {
	Create(request CreateRequest)
	Get(request GetRequest) Response
	Update(request UpdateRequest)
}
type Service struct {
	repo     Repository
	Validate *validator.Validate
}

func NewService(repo Repository, validate *validator.Validate) *Service {
	return &Service{
		repo:     repo,
		Validate: validate,
	}
}
func (s *Service) Create(request CreateRequest) error {
	err := s.Validate.RegisterValidation("isContactValid", common_go.ValidateContactNo)
	if err != nil {
		fmt.Println("Error in validation:", err.Error())
	}
	err = s.Validate.Struct(request)
	if err != nil {
		panic("Invalid Entries")
	}
	id, err := gonanoid.Generate(`"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"`, 5)
	patientModel := Patient{
		ID:        id,
		Name:      request.Name,
		ContactNo: request.ContactNo,
		Address:   request.Address,
		DoctorID:  request.DoctorID,
	}
	err = s.repo.Create(&patientModel)
	return err
}

func (s *Service) Get(request GetRequest) (Response, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		panic("Invalid ID")
	}
	pat, err2 := s.repo.Get(request.ID)
	response := Response{ID: pat.ID, CreatedAt: pat.CreatedAt, UpdatedAt: pat.UpdatedAt, Name: pat.Name, ContactNo: pat.ContactNo, Address: pat.Address, DoctorID: pat.DoctorID}
	return response, err2
}
func (s *Service) Update(id string, request UpdateRequest) error {
	err := s.Validate.RegisterValidation("isContactValid", common_go.ValidateContactNo)
	if err != nil {
		fmt.Println("Error in validation:", err.Error())
	}
	err = s.Validate.Struct(request)
	if err != nil {
		panic("Invalid Entries")
	}
	pat, err1 := s.repo.Get(id)
	pat.ContactNo = request.ContactNo
	pat.Address = request.Address
	pat.DoctorID = request.DoctorID
	if err1 != nil {
		return err
	}
	err = s.repo.Update(pat)
	return err
}
