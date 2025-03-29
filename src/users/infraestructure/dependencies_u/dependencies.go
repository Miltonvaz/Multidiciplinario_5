package dependencies_u

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/users/application"
	db "Multidiciplinario/src/users/infraestructure/adapters"
	"Multidiciplinario/src/users/infraestructure/controllers"
)

func Init() (
	*controllers.CreateUserController,
	*controllers.ViewUserController,
	*controllers.EditUserController,
	*controllers.DeleteUserController,
	*controllers.ViewUserByIdController,
	*controllers.AuthController,
	error,
) {

	pool := core.GetDBPool()
	ps := db.NewMySQL(pool.DB)

	createClient := application.NewCreateUser(ps)
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
