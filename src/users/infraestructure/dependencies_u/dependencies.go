package dependencies_u

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/users/application"
	"Multidiciplinario/src/users/application/repositories"
	"Multidiciplinario/src/users/infraestructure/adapters"
	"Multidiciplinario/src/users/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.CreateUserController,
	*controllers.ViewUserController,
	*controllers.EditUserController,
	*controllers.DeleteUserController,
	*controllers.ViewUserByIdController,
	*controllers.AuthController,
	error,
) {

	ps := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	createClient := application.NewCreateUser(ps, serviceNotification)
	viewClient := application.NewListUser(ps)
	editClient := application.NewEditUser(ps)
	deleteClient := application.NewDeleteUser(ps)
	viewClientById := application.NewUserById(ps)
	authService := application.NewAuthService(ps)

	authController := controllers.NewAuthController(authService)
	createClientController := controllers.NewCreateUserController(*createClient)
	viewClientController := controllers.NewViewUserController(*viewClient)
	editClientController := controllers.NewEditUserController(*editClient)
	deleteClientController := controllers.NewDeleteUserController(*deleteClient)
	viewClientByIdController := controllers.NewViewUserByIdController(*viewClientById)

	return createClientController, viewClientController, editClientController, deleteClientController, viewClientByIdController, authController, nil
}
