package doctor

import (
	"fmt"
	"hospital_management_service/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

func (ctrl *Controller) Create(c *gin.Context) {
	var doc CreateRequest
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.Create(doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}
	c.JSON(http.StatusOK, doc)
}

func (ctrl *Controller) Get(c *gin.Context) {
	id := c.Param(constants.ID)
	fmt.Printf(id)
	doc, err := ctrl.Service.Get(GetRequest{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	c.JSON(http.StatusOK, doc)
}

func (ctrl *Controller) Update(c *gin.Context) {
	var doc UpdateRequest
	id := c.Param(constants.ID)
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.Service.Update(id, doc)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, doc)
}
func (ctrl *Controller) FetchPatientsByDoctorID(c *gin.Context) {
	id := c.Param(constants.ID)
	pat, err := ctrl.Service.FetchPatientsByDoctorID(FetchPatientsByDoctorID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	c.JSON(http.StatusOK, pat)
}
