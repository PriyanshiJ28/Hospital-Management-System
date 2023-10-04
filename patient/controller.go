package patient

import (
	"github.com/gin-gonic/gin"
	"hospital_management_service/constants"
	"net/http"
)

type Controller struct {
	Service Service
}

func NewController(service Service) *Controller {
	return &Controller{Service: service}
}

func (ctrl *Controller) Create(c *gin.Context) {
	var pat CreateRequest
	if err := c.ShouldBindJSON(&pat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.Create(pat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}
	c.JSON(http.StatusCreated, pat)
}

func (ctrl *Controller) Get(c *gin.Context) {
	id := c.Param(constants.ID)
	pat, err := ctrl.Service.Get(GetRequest{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, pat)
}

func (ctrl *Controller) Update(c *gin.Context) {
	var pat UpdateRequest
	id := c.Param(constants.ID)
	if err := c.ShouldBindJSON(&pat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.Service.Update(id, pat)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, pat)
}
