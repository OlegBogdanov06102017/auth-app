package handler

import (
	"net/http"

	authapp "github.com/OlegBogdanov06102017/auth-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(c *gin.Context) {
	var input authapp.Customer

	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateCustomer(input)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) singIn(c *gin.Context) {

}
