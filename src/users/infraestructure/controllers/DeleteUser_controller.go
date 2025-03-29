package controllers

import (
	"Multidiciplinario/src/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	usecase application.DeleteUser
}

func NewDeleteUserController(usecase application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{usecase: usecase}
}

func (dc_c *DeleteUserController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := dc_c.usecase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
