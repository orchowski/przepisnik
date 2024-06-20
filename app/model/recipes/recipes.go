package recipes

import (
	model "fit.synapse/przepisnik/app/model/dto"
	"github.com/google/uuid"
)

type Facade interface {
	Create(id *uuid.UUID, title string) (*model.Recipe, error)
	Delete(id uuid.UUID) error
	AddStage(name string, description string, recipeId uuid.UUID) (*model.Recipe, error)
	AddIngredientToStage(ingredientId uuid.UUID, stageId uuid.UUID) error
	AddPhoto(photoPath string, recipeId uuid.UUID, stage string, stepIndex *byte) error
	AddStep(recipeId uuid.UUID, stage string, description string) (*model.Recipe, error)
}

type recipes struct {
}

func InitializeRecipesModule() Facade {
	return &recipes{}
}

func (r *recipes) Create(id *uuid.UUID, title string) (*model.Recipe, error) {
	return nil, nil
}

func (r *recipes) Delete(id uuid.UUID) error {
	return nil
}

func (r *recipes) AddStage(name string, description string, recipeId uuid.UUID) (*model.Recipe, error) {
	return nil, nil
}

func (r *recipes) AddIngredientToStage(ingredientId uuid.UUID, stageId uuid.UUID) error {
	return nil
}

func (r *recipes) AddPhoto(photoPath string, recipeId uuid.UUID, stage string, stepIndex *byte) error {
	return nil
}

func (r *recipes) AddStep(recipeId uuid.UUID, stage string, description string) (*model.Recipe, error) {
	return nil, nil
}
