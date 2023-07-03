package api

import (
	"context"
	"net/http"
	"strconv"
	pbhospital "testing/pb/hospital"

	"github.com/gin-gonic/gin"
)

type HospitalHandler struct {
	client pbhospital.HospitalServiceClient
}

func NewHospitalHandler(client pbhospital.HospitalServiceClient) *HospitalHandler {
	return &HospitalHandler{client: client}
}

// Create Hospital
func (h *HospitalHandler) CreateHospital(c *gin.Context) {
	// make request to create hospital
	req := &pbhospital.Hospital{
		Hospitalcode: c.PostForm("hospitalcode"),
		Name:         c.PostForm("name"),
		Address:      c.PostForm("address"),
	}

	// get response from server
	res, err := h.client.CreateHospital(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, res)
}

// Get All Hospital
func (h *HospitalHandler) GetHospital(c *gin.Context) {
	// make request to get all hospital
	req := &pbhospital.HospitalRequest{}

	// get response from server
	res, err := h.client.GetHospital(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, res)
}

// Get Patient By ID
func (h *HospitalHandler) GetPatientById(c *gin.Context) {
	// make request to get patient by id
	req := &pbhospital.PatientId{
		IdentityId: c.Param("id"),
	}

	// get response from server
	res, err := h.client.GetPatientById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, res)
}

// Add Patient
func (h *HospitalHandler) AddPatient(c *gin.Context) {
	// make request to add patient
	age, err := strconv.ParseInt(c.PostForm("age"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age"})
		return
	}
	req := &pbhospital.Patient{
		IdentityId: c.PostForm("identity_id"),
		Name:       c.PostForm("name"),
		Age:        int32(age),
		Disease:    c.PostForm("disease"),
	}

	// get response from server
	res, err := h.client.AddPatient(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, res)
}

// get patients
func (h *HospitalHandler) GetPatients(c *gin.Context) {
	// make request to get patients
	req := &pbhospital.PatientRequest{}

	// get response from server
	res, err := h.client.GetPatients(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return response
	c.JSON(http.StatusOK, res)
}

