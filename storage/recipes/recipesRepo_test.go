package storage

import (
	"testing"

	model "fit.synapse/przepisnik/app/model/dto"
	"fit.synapse/przepisnik/utility"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var ingredientDefinitions map[string]model.IngredientDefinition
var authorId uuid.UUID

func newStorage() *recipesRepo {
	authorId = *utility.Newuuid()
	ingredientStorage := newIngredientsInMemoryStorage()
	recipeStorage := NewRecipesInMemoryStorage(ingredientStorage)

	ingredientDefinitions = make(map[string]model.IngredientDefinition)

	ingredientDefinitions["Tomato"] = model.IngredientDefinition{
		Id:   utility.Newuuid(),
		Name: "Tomato",
		Kcal: &model.KcalUnit{
			Gr100:    30,
			UnitGr:   utility.AddressOf(100),
			UnitName: utility.AddressOf("Piece"),
		},
	}
	ingredientStorage.Upsert(ingredientDefinitions["Tomato"])

	ingredientDefinitions["Pork"] = model.IngredientDefinition{
		Id:   utility.Newuuid(),
		Name: "Pork",
		Kcal: &model.KcalUnit{
			Gr100: 333,
		},
	}
	ingredientStorage.Upsert(ingredientDefinitions["Pork"])

	ingredientDefinitions["Water"] = model.IngredientDefinition{
		Id:   utility.Newuuid(),
		Name: "Water",
		Kcal: &model.KcalUnit{
			Gr100:    0,
			UnitGr:   utility.AddressOf(250),
			UnitName: utility.AddressOf("Cup"),
		},
	}
	ingredientStorage.Upsert(ingredientDefinitions["Water"])

	ingredientDefinitions["Flour"] = model.IngredientDefinition{
		Id:   utility.Newuuid(),
		Name: "Flour",
		Kcal: &model.KcalUnit{
			Gr100:    400,
			UnitGr:   utility.AddressOf(250),
			UnitName: utility.AddressOf("Cup"),
		},
	}
	ingredientStorage.Upsert(ingredientDefinitions["Flour"])

	return recipeStorage
}

func TestNewRecipesStorage(t *testing.T) {
	storage := newStorage()
	if storage == nil {
		t.Fatal("Expected non-nil storage")
	}

	if storage.recipes == nil {
		t.Fatal("Expected non-nil map of users")
	}
}

func TestCreateNewRecipe(t *testing.T) {
	storage := newStorage()
	newRecipe := fullRecipeExample(nil)
	recipeId, err := storage.Upsert(newRecipe)
	newRecipe.Id = recipeId

	if err != nil {
		t.Fatal("Expected no errors while creating recipe, got: " + err.Error())
	}
	if recipeId == nil {
		t.Fatal("recipe id should not be nil")
	}

	recipeResult := storage.Get(*recipeId)

	assert.Equal(t, newRecipe, recipeResult, "Recipe after retrieval should be equal")
}

func TestUpdateRecipe(t *testing.T) {
	storage := newStorage()
	newRecipe := fullRecipeExample(nil)
	recipeId, err := storage.Upsert(newRecipe)
	newRecipe.Id = recipeId

	if err != nil {
		t.Fatal("Expected no errors while creating recipe, got: " + err.Error())
	}
	if recipeId == nil {
		t.Fatal("recipe id should not be nil")
	}

	recipeResult := storage.Get(*recipeId)

	assert.Equal(t, newRecipe, recipeResult, "New recipe after fetch should be the same")

	updatedRecipe := newRecipe
	updatedRecipe.Stages["Sauce"] = []*model.Ingredient{}
	recipeId, err = storage.Upsert(updatedRecipe)

	if err != nil {
		t.Fatal("Expected no errors while creating recipe, got: " + err.Error())
	}
	if recipeId == nil {
		t.Fatal("recipe id should not be nil")
	}

	recipeResult = storage.Get(*recipeId)

	assert.Equal(t, newRecipe, recipeResult, "Updated recipe after fetch should be the same")
}

func TestDeleteRecipe(t *testing.T) {
	storage := newStorage()
	newRecipe := fullRecipeExample(nil)
	recipeId, err := storage.Upsert(newRecipe)
	newRecipe.Id = recipeId

	if err != nil {
		t.Fatal("Expected no errors while creating recipe, got: " + err.Error())
	}
	if recipeId == nil {
		t.Fatal("recipe id should not be nil")
	}

	recipeResult := storage.Get(*recipeId)

	assert.Equal(t, newRecipe, recipeResult, "New recipe after fetch should be the same")

	err = storage.Delete(*recipeId)
	if err != nil {
		t.Fatal("Error occured while deleting: " + err.Error())
	}
	recipeResult = storage.Get(*recipeId)

	assert.Nil(t, recipeResult, "Recipe after deletion should not be present")
}

func fullRecipeExample(id *uuid.UUID) *model.Recipe {
	pork := ingredientDefinitions["Pork"]
	tomato := ingredientDefinitions["Tomato"]
	water := ingredientDefinitions["Water"]
	flour := ingredientDefinitions["Flour"]

	return &model.Recipe{
		Id:   id,
		Name: "Spaghetti bolognesse",
		Stages: map[string][]*model.Ingredient{
			"Sauce": {
				{
					Id:   tomato.Id,
					Name: tomato.Name,
					Amount: model.IngredientAmount{
						Gr:    400,
						Units: utility.AddressOf(4.0),
						Kcal:  tomato.Kcal,
					},
				},
				{
					Id:   pork.Id,
					Name: pork.Name,
					Amount: model.IngredientAmount{
						Gr:    400,
						Units: new(float64),
						Kcal:  pork.Kcal,
					},
				},
			},
			"Pasta": {
				{
					Id:   water.Id,
					Name: water.Name,
					Amount: model.IngredientAmount{
						Gr:    500,
						Units: utility.AddressOf(1.0),
						Kcal:  water.Kcal,
					},
				},
				{
					Id:   flour.Id,
					Name: flour.Name,
					Amount: model.IngredientAmount{
						Gr:    250,
						Units: utility.AddressOf(1.0),
						Kcal:  flour.Kcal,
					},
				},
			},
		},
		Steps: map[string][]*model.Step{
			"Sauce": {
				{
					Description: "do things",
					PhotoURLs:   []string{"/path1"},
				},
				{
					Description: "do things",
				},
			},
		},
		PhotoURLs: []string{"/path2"},
		Author:    authorId,
		Portions:  4,
		Tags:      []string{"Pasta", "Italy"},
	}

}
