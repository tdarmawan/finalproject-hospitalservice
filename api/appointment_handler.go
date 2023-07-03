package api

import (
	"context"
	"net/http"
	"strconv"
	pbappointment "testing/pb/appointment"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	client pbappointment.AppointmentServiceClient
}

func NewAppointmentHandler(client pbappointment.AppointmentServiceClient) *AppointmentHandler {
	return &AppointmentHandler{client: client}
}

func (h *AppointmentHandler) GetAppointmentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	req := &pbappointment.AppointmentId{
		//appointmentcode must be string
		AppointmentCode: strconv.Itoa(id),
	}
	res, err := h.client.ReadAppointmentById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	req := &pbappointment.CreateAppointmentRequest{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.CreateAppointment(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) GetAllAppointment(c *gin.Context) {
	req := &pbappointment.Empty{}

	res, err := h.client.ReadAppointmentAll(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *AppointmentHandler) EmailAppointment(c *gin.Context) {
	req := &pbappointment.EmailAppointmentRequest{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.EmailAppointment(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
