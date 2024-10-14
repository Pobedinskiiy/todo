package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.newErrorResponse(c, http.StatusUnauthorized, "missing authorization header")
		return
	}

	headerParts := strings.SplitN(header, " ", 2)
	if len(headerParts) != 2 {
		h.newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		h.newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return 0, errors.New("user not found")
	}

	idInt, ok := id.(int)
	if !ok {
		h.newErrorResponse(c, http.StatusInternalServerError, "user is not of invalid type")
		return 0, errors.New("user is not of invalid type")
	}

	return idInt, nil
}
