package common_go

import (
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

func ValidateContactNo(fl validator.FieldLevel) bool {
	contactNo := fl.Field().String()
	matched, _ := regexp.MatchString(`^\d{10}$`, contactNo)
	return matched
}

//func PatientRequestToResponse(pat *patient.Patient) patient.Response {
//	response := patient.Response{
//		ID:        pat.ID,
//		CreatedAt: pat.CreatedAt,
//		UpdatedAt: pat.UpdatedAt,
//		Name:      pat.Name,
//		ContactNo: pat.ContactNo,
//		Address:   pat.Address,
//		DoctorID:  pat.DoctorID,
//	}
//	return response
//}
//func DoctorRequestToResponse(doc *doctor.Doctor) doctor.Response {
//	response := doctor.Response{
//		ID:        doc.ID,
//		CreatedAt: doc.CreatedAt,
//		UpdatedAt: doc.UpdatedAt,
//		Name:      doc.Name,
//		ContactNo: doc.ContactNo,
//		Patients:  doc.Patients,
//	}
//	return response
//}
