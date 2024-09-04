package storage

import (
	model "fit.synapse/przepisnik/app/model/dto"
	"github.com/google/uuid"
)

func MapRecipeDto(id uuid.UUID, recipeDto model.Recipe) *recipe {
	mapStages := func(stagesDto map[model.StageName][]*model.Ingredient) map[string][]*recipeIngredient {
		stages := make(map[string][]*recipeIngredient)
		mapIngredients := func(ingredients []*model.Ingredient) []*recipeIngredient {
			recipeIngredients := []*recipeIngredient{}
			for _, ingredient := range ingredients {
				recipeIngredients = append(recipeIngredients, &recipeIngredient{
					Id:          ingredient.Id,
					AmountGr:    ingredient.Amount.Gr,
					AmountUnits: ingredient.Amount.Units,
				})
			}
			return recipeIngredients
		}

		for stgName, ingredients := range stagesDto {
			stages[stgName] = mapIngredients(ingredients)
		}

		return stages
	}

	mapSteps := func(stagesDto map[model.StageName][]*model.Step) map[string][]*step {
		resultSteps := make(map[string][]*step)
		mapStepList := func(steps []*model.Step) []*step {
			stepList := []*step{}
			for _, singleStep := range steps {
				stepList = append(stepList, &step{
					Description: singleStep.Description,
					PhotoURLs:   singleStep.PhotoURLs,
				})
			}
			return stepList
		}

		for stgName, steps := range stagesDto {
			resultSteps[stgName] = mapStepList(steps)
		}

		return resultSteps
	}

	recipe := &recipe{
		Id:        id,
		Name:      recipeDto.Name,
		Stages:    mapStages(recipeDto.Stages),
		Steps:     mapSteps(recipeDto.Steps),
		PhotoURLs: recipeDto.PhotoURLs,
		Author:    recipeDto.Author,
		Portions:  recipeDto.Portions,
		Tags:      recipeDto.Tags,
		Calories:  recipeDto.Calories,
		Rating:    recipeDto.Rating,
	}
	return recipe
}

// todo: in progress
func MapRecipeDb(id uuid.UUID, recipeDb recipe, allRecipeIngredients []*ingredient) *model.Recipe {
	mapStages := func(stagesDto map[model.StageName][]*model.Ingredient) map[string][]ingredientId {
		stages := make(map[string][]uuid.UUID)
		mapIngredients := func(ingredients []*model.Ingredient) []uuid.UUID {
			ingredientIds := []uuid.UUID{}
			for _, ingredient := range ingredients {
				ingredientIds = append(ingredientIds, ingredient.Id)
			}
			return ingredientIds
		}

		for stgName, ingredients := range stagesDto {
			stages[stgName] = mapIngredients(ingredients)
		}

		return stages
	}

	mapSteps := func(stagesDto map[model.StageName][]*model.Step) map[string][]*step {
		resultSteps := make(map[string][]*step)
		mapStepList := func(steps []*model.Step) []*step {
			stepList := []*step{}
			for _, singleStep := range steps {
				stepList = append(stepList, &step{
					Description: singleStep.Description,
					PhotoURLs:   singleStep.PhotoURLs,
				})
			}
			return stepList
		}

		for stgName, steps := range stagesDto {
			resultSteps[stgName] = mapStepList(steps)
		}

		return resultSteps
	}

	recipe := &recipe{
		Id:        id,
		Name:      recipeDb.Name,
		Stages:    mapStages(recipeDb.Stages),
		Steps:     mapSteps(recipeDb.Steps),
		PhotoURLs: recipeDb.PhotoURLs,
		Author:    recipeDb.Author,
		Portions:  recipeDb.Portions,
		Tags:      recipeDb.Tags,
		Calories:  recipeDb.Calories,
		Rating:    recipeDb.Rating,
	}
	return recipe
}
