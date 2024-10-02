package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "os/user"
	"todo"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) singIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
