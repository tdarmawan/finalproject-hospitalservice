package api

import (
	"context"
	"net/http"
	pbuser "testing/pb/user"

	"github.com/gin-gonic/gin"
)

// UserHandler is a struct that holds the UserDatakuClient.
type UserHandler struct {
	client pbuser.UserDatakuClient
}

// NewUserHandler returns a new UserHandler instance.
func NewUserHandler(client pbuser.UserDatakuClient) *UserHandler {
	return &UserHandler{
		client: client,
	}
}

// GetAllUserData handles HTTP POST requests to the /v1/users endpoint.
func (h *UserHandler) GetAllUserData(c *gin.Context) {

	req := &pbuser.GetAllUserRequest{}
	// fmt.Println(req, "ss")
	// fmt.Println(c.BindJSON(req), "asa")
	// if err := c.BindJSON(req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	resp, err := h.client.GetAllData(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserData handles HTTP GET requests to the /v1/users/:id endpoint.
func (h *UserHandler) GetUserData(c *gin.Context) {
	id := c.Param("id")
	req := &pbuser.GetUserRequest{
		UserId: id,
	}

	resp, err := h.client.GetUserData(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddUserData handles HTTP POST requests to the /v1/users/add endpoint.
func (h *UserHandler) AddUserData(c *gin.Context) {
	req := &pbuser.AddUserRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.AddUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUserData handles HTTP PUT requests to the /v1/users/:id endpoint.
func (h *UserHandler) UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	req := &pbuser.UpdateUserRequest{
		Id: id,
	}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.UpdateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AuthUser handles HTTP POST requests to the /v1/auth endpoint.
func (h *UserHandler) AuthUser(c *gin.Context) {
	req := &pbuser.AuthRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.AuthUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GenerateToken handles HTTP POST requests to the /v1/token endpoint.
func (h *UserHandler) GenerateToken(c *gin.Context) {
	req := &pbuser.GenerateTokenRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.client.GenerateToken(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// StartScheduledJob handles HTTP POST requests to the /v1/jobs/start endpoint.
// func (h *UserHandler) StartScheduledJob(c *gin.Context) {
// 	req := &pb.StartScheduledJobRequest{}
// 	if err := c.BindJSON(req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	resp, err := h.client.StartScheduledJob(context.Background(), req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }
