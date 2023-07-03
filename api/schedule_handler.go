package api

import (
	"context"
	"net/http"
	"strconv"
	pbschedule "testing/pb/schedule"

	"github.com/gin-gonic/gin"
)

// ScheduleHandler is a struct.
type ScheduleHandler struct {
	client pbschedule.ScheduleAvailabilityDoctorServiceClient
}

// NewScheduleHandler is a function.
func NewScheduleHandler(client pbschedule.ScheduleAvailabilityDoctorServiceClient) *ScheduleHandler {
	return &ScheduleHandler{client: client}
}

func (h *ScheduleHandler) GetAllScheduleByDoctorId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	req := &pbschedule.ScheduleAvailabilityDoctorRequestByDoctorId{
		DoctorId: int32(id),
	}

	res, err := h.client.ReadScheduleByDoctorId(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ScheduleHandler) GetScheduleByDay(c *gin.Context) {
	req := &pbschedule.ScheduleAvailabilityDoctorRequestByDay{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.ReadScheduleByDay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ScheduleHandler) GetAllSchedule(c *gin.Context) {
	req := &pbschedule.ScheduleAvailabilityDoctorRequest{}

	res, err := h.client.Read(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	req := &pbschedule.ScheduleAvailabilityDoctorRequest{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.Create(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	req := &pbschedule.ScheduleAvailabilityDoctorRequest{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.Update(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	req := &pbschedule.ScheduleAvailabilityDoctorRequest{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.Delete(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
