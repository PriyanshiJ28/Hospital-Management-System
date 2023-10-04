package main

import (
	"gopkg.in/go-playground/validator.v9"
	"hospital_management_service/config"
	doctorController "hospital_management_service/doctor"
	patientController "hospital_management_service/patient"
	"hospital_management_service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb(DBConfig config.DatabaseConfig, servConfig config.ServiceConfig) {
	db = config.DatabaseConnection(DBConfig, servConfig)
	err := db.AutoMigrate(&doctorController.Doctor{}, &patientController.Patient{})
	if err != nil {
		return
	}
}

func initRouter() {
	router := gin.Default()

	validate := validator.New()
	doctorRepo := doctorController.NewRepository(db)
	patientRepo := patientController.NewRepository(db)

	doctorServ := doctorController.NewService(*doctorRepo, validate)
	patientServ := patientController.NewService(*patientRepo, validate)

	doctorCtrl := doctorController.NewController(*doctorServ)
	patientCtrl := patientController.NewController(*patientServ)

	routes.SetupRoutes(router, doctorCtrl, patientCtrl)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func main() {
	DBConfig, servConfig := config.LoadConfig()
	initDb(DBConfig, servConfig)
	initRouter()
}
