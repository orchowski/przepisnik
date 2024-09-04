package storage

import (
	"errors"

	model "fit.synapse/przepisnik/app/model/dto"
	commons "fit.synapse/przepisnik/commons"
	utility "fit.synapse/przepisnik/utility"
	"github.com/google/uuid"
)

type recipesRepo struct {
	recipes map[uuid.UUID]*recipe
}

func NewRecipesPersistentStorage(basePath string) *recipesRepo {
	return NewRecipesInMemoryStorage()
}

func NewRecipesInMemoryStorage() *recipesRepo {
	return &recipesRepo{recipes: make(map[uuid.UUID]*recipe)}
}

func (rp *recipesRepo) Upsert(recipeDto model.Recipe) (*uuid.UUID, error) {
	var id uuid.UUID

	if recipeDto.Id == nil {
		id = uuid.New()
	} else {
		id = *recipeDto.Id
	}
	recipe := MapRecipe(id, recipeDto)

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

	return &model.Recipe{
		Id: &recipeFound.Id,
	}
}

func (rp *recipesRepo) GetAll() []*model.Recipe {
	var allUsers []*model.Recipe
	for _, user := range rp.users {
		allUsers = append(allUsers, &model.Recipe{})
	}
	return allUsers
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
