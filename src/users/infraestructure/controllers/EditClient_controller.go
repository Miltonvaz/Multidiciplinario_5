package controllers

import (
	"Multidiciplinario/src/users/application"
	"Multidiciplinario/src/users/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditUserController struct {
	usecase application.EditUser
}

func NewEditUserController(usecase application.EditUser) *EditUserController {
	return &EditUserController{usecase: usecase}
}

func (ed_c *EditUserController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request entities.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	request.ID = int32(id)

	if err := ed_c.usecase.Execute(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.Password = ""

	c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully", "Client": request})
}
