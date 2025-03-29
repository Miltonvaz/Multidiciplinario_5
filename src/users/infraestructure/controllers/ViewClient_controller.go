package controllers

import (
	"Multidiciplinario/src/users/application"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewUserController struct {
	useCase application.ViewUser
}

func NewViewUserController(useCase application.ViewUser) *ViewUserController {
	return &ViewUserController{useCase: useCase}
}

func (cc_c *ViewUserController) Execute(c *gin.Context) {
	users, err := cc_c.useCase.Execute()
	if err != nil {
		// Log para capturar el error de manera más detallada
		fmt.Printf("Error retrieving users: %v\n", err)
		// Enviar una respuesta con detalles del error
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to retrieve clients: %v", err)})
		return
	}

	// Enviar los usuarios en caso de éxito
	c.JSON(http.StatusOK, gin.H{"users": users})
}
