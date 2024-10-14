package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo"
)

func (h *Handler) createList(c *gin.Context) {
	id, err := h.getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.ShouldBindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.service.TodoList.Create(id, input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
