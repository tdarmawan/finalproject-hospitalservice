package api

import (
	"context"
	"net/http"
	"strconv"
	pbdoctor "testing/pb/doctor"

	"github.com/gin-gonic/gin"
)

// DoctorHandle is a struct.
type DoctorHandler struct {
	client pbdoctor.DoctorServiceClient
}

// NewDoctorHandler is a function.
func NewDoctorHandler(client pbdoctor.DoctorServiceClient) *DoctorHandler {
	return &DoctorHandler{client: client}
}

// GetDoctor is a function.
func (h *DoctorHandler) GetDoctorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	req := &pbdoctor.DoctorId{
		DoctorId: int32(id),
	}
	res, err := h.client.GetDoctorById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CreateDoctor is a function.
func (h *DoctorHandler) CreateDoctor(c *gin.Context) {
	//make request from &pbdoctor.DoctorRequest that accept json input
	req := &pbdoctor.DoctorRequest{}

	//bind json input to request
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.CreateDoctor(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateDoctor is a function.
func (h *DoctorHandler) UpdateDoctor(c *gin.Context) {
	req := &pbdoctor.DoctorRequest{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.UpdateDoctor(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteDoctor is a function.
func (h *DoctorHandler) DeleteDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	req := &pbdoctor.DoctorId{
		DoctorId: int32(id),
	}
	res, err := h.client.DeleteDoctor(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetDoctorByHospital is a function.
func (h *DoctorHandler) GetDoctorByHospitalId(c *gin.Context) {
	hospitalID, err := strconv.Atoi(c.Param("hospital_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hospital ID"})
		return
	}
	req := &pbdoctor.HospitalId{
		HospitalId: int32(hospitalID),
	}
	res, err := h.client.GetDoctorByHospitalId(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllDoctor is a function.
func (h *DoctorHandler) GetAllDoctors(c *gin.Context) {
	req := &pbdoctor.Empty{}
	res, err := h.client.GetAllDoctors(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
