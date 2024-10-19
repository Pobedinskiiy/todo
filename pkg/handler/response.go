package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	h.log.Warn(fmt.Sprintf("Failed to response, error: %s", message))
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
