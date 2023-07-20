package utils

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the error response format.
type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

// NewErrorResponse is a helper function to send an error response to the client with the given status code and message.
// It also logs the error message.
// @Summary Send an error response to the client.
// @Tags Error
// @Param statusCode query int true "Status code of the error response"
// @Param message query string true "Error message"
// @Produce json
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
