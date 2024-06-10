package recipes

import "fit.synapse/przepisnik/app/model/dto"

type Recipes interface {
	CreateNew() (*model.Recipe, *error)
	Delete() (*model.Recipe, *error)
	AddStage() (*model.Recipe, *error)
	AddIngredient() (*model.Recipe, *error)
	AddStep() (*model.Recipe, *error)
	AddPhoto()
}
