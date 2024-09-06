package storage

import (
	"errors"

	model "fit.synapse/przepisnik/app/model/dto"
	commons "fit.synapse/przepisnik/commons"
	utility "fit.synapse/przepisnik/utility"
	"github.com/google/uuid"
)

type recipesRepo struct {
	recipes         map[uuid.UUID]*recipe
	ingredientsRepo *ingredientsRepo
}

func NewRecipesPersistentStorage(basePath string) *recipesRepo {
	return NewRecipesInMemoryStorage(newIngredientsInMemoryStorage())
}

func NewRecipesInMemoryStorage(ingredientsRepository *ingredientsRepo) *recipesRepo {
	return &recipesRepo{recipes: make(map[uuid.UUID]*recipe), ingredientsRepo: ingredientsRepository}
}

func (rp *recipesRepo) Upsert(recipeDto model.Recipe) (*uuid.UUID, error) {
	var id uuid.UUID

	if recipeDto.Id == nil {
		id = uuid.New()
	} else {
		id = *recipeDto.Id
	}
	recipe := MapRecipeDto(id, recipeDto)

	if validationError := isValid(recipe); validationError != nil {
		return nil, validationError
	}

	rp.recipes[id] = recipe
	return &id, nil
}

func (rp *recipesRepo) Get(id uuid.UUID) *model.Recipe {
	recipeFound := rp.recipes[id]
	if recipeFound == nil || recipeFound.Deleted {
		return nil
	}

	return MapRecipeDb(recipeFound, rp.ingredientsRepo.getAllForRecipe(recipeFound))
}

func (rp *recipesRepo) Delete(id uuid.UUID) error {
	recipe := rp.recipes[id]
	if recipe == nil {
		return commons.RecipeNotFound
	}
	recipe.Deleted = true
	return nil
}
func isValid(recipe *recipe) error {
	if len(recipe.Name) == 0 {
		return errors.New("name must be provided")
	}
	stagesKeys := make([]string, 0, len(recipe.Stages))
	for k := range recipe.Stages {
		stagesKeys = append(stagesKeys, k)
	}
	stepsKeys := make([]string, 0, len(recipe.Steps))
	for k := range recipe.Steps {
		stepsKeys = append(stepsKeys, k)
	}
	utility.IsSubset(stepsKeys, stagesKeys)

	return nil
}
