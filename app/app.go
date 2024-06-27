package app

import (
	recipes "fit.synapse/przepisnik/app/model/recipes"
	"fit.synapse/przepisnik/app/model/users"
)

type Application struct {
	Recipes recipes.Facade
	Users   users.Facade
}

func NewApp(basePath string) *Application {
	return &Application{
		Recipes: recipes.InitializeRecipesModule(),
		Users:   users.InitializeUsersModule(basePath),
	}
}
