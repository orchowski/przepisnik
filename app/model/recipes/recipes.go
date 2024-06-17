package recipes

import (
	"fit.synapse/przepisnik/app/model/dto"
	"github.com/google/uuid"
)

type Recipes interface {
	Create(id *uuid.UUID, title string) (*model.Recipe, error)
	Delete(id uuid.UUID) error
	AddStage(description string, recipeId uuid.UUID) (*model.Recipe, error)
	AddIngredientToStage(ingredientId uuid.UUID, stageId uuid.UUID) error
	AddPhoto(photoPath string, recipeId uuid.UUID, stage string, stepIndex *byte) error
	AddStep(recipeId uuid.UUID, stage string, description string) (*model.Recipe, error)
}
