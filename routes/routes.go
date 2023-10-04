package routes

import (
	"github.com/gin-gonic/gin"
	dController "hospital_management_service/doctor"
	pController "hospital_management_service/patient"
)

func SetupRoutes(router *gin.Engine, doctorController *dController.Controller, patientController *pController.Controller) {
	router.POST("/doctor/", doctorController.Create)
	router.GET("/doctor/:id", doctorController.Get)
	router.PATCH("/doctor/:id", doctorController.Update)

	router.POST("/patient/", patientController.Create)
	router.GET("/patient/:id", patientController.Get)
	router.PATCH("/patient/:id", patientController.Update)

	router.GET("/fetchPatientByDoctorId/:id", doctorController.FetchPatientsByDoctorID)
}
