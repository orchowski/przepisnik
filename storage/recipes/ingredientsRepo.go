package storage

import (
	model "fit.synapse/przepisnik/app/model/dto"
	"github.com/google/uuid"
)

type ingredientsRepo struct {
	ingredients map[uuid.UUID]*ingredient
}

func newIngredientsPersistentStorage(basePath string) *ingredientsRepo {
	return newIngredientsInMemoryStorage()
}

func newIngredientsInMemoryStorage() *ingredientsRepo {
	return &ingredientsRepo{ingredients: make(map[uuid.UUID]*ingredient)}
}

func (ir *ingredientsRepo) Upsert(ingredientDef model.IngredientDefinition) (*uuid.UUID, error) {
	var id uuid.UUID

	if ingredientDef.Id == nil {
		id = uuid.New()
	} else {
		id = *ingredientDef.Id
	}

	//	if validationError := isValid(recipe); validationError != nil {
	//		return nil, validationError
	//	}

	ir.ingredients[id] = &ingredient{
		Id:   id,
		Name: ingredientDef.Name,
		Kcal: &kcalUnit{
			Gr100:    ingredientDef.Kcal.Gr100,
			UnitGr:   ingredientDef.Kcal.UnitGr,
			UnitName: ingredientDef.Kcal.UnitName,
		},
	}
	return &id, nil
}

func (ir *ingredientsRepo) Get(id uuid.UUID) *model.IngredientDefinition {
	ingredientFound := ir.ingredients[id]
	if ingredientFound == nil {
		return nil
	}

	return &model.IngredientDefinition{
		Id:   &ingredientFound.Id,
		Name: ingredientFound.Name,
		Kcal: &model.KcalUnit{
			Gr100:    ingredientFound.Kcal.Gr100,
			UnitGr:   ingredientFound.Kcal.UnitGr,
			UnitName: ingredientFound.Kcal.UnitName,
		},
	}
}

func (ir *ingredientsRepo) getAllForRecipe(recipe *recipe) map[uuid.UUID]*ingredient {
	ingredients := make(map[uuid.UUID]*ingredient)
	for _, values := range recipe.Stages {
		for _, each := range values {
			ingredients[each.Id] = ir.ingredients[each.Id]
		}
	}

	return ingredients
}
