package storage

import (
	model "fit.synapse/przepisnik/app/model/dto"
	"github.com/google/uuid"
)

func MapRecipeDto(id uuid.UUID, recipeDto *model.Recipe) *recipe {
	mapStages := func(stagesDto map[model.StageName][]*model.Ingredient) map[string][]*recipeIngredient {
		stages := make(map[string][]*recipeIngredient)
		mapIngredients := func(ingredients []*model.Ingredient) []*recipeIngredient {
			recipeIngredients := []*recipeIngredient{}
			for _, ingredient := range ingredients {
				recipeIngredients = append(recipeIngredients, &recipeIngredient{
					Id:          *ingredient.Id,
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

func MapRecipeDb(recipeDb *recipe, allRecipeIngredientDefinitions map[uuid.UUID]*ingredient) *model.Recipe {
	mapStages := func(stagesDto map[model.StageName][]*recipeIngredient) map[string][]*model.Ingredient {
		stages := make(map[string][]*model.Ingredient)
		mapIngredients := func(ingredients []*recipeIngredient) []*model.Ingredient {
			ingredientsMapped := []*model.Ingredient{}
			for _, ingredient := range ingredients {
				fullIngredient := allRecipeIngredientDefinitions[ingredient.Id]
				if fullIngredient == nil {
					ingredientsMapped = append(ingredientsMapped, &model.Ingredient{
						Name: "Ingredient deleted",
					})
					continue
				}
				ingredientsMapped = append(ingredientsMapped, &model.Ingredient{
					Id:   &ingredient.Id,
					Name: fullIngredient.Name,
					Amount: model.IngredientAmount{
						Gr:    ingredient.AmountGr,
						Units: ingredient.AmountUnits,
						Kcal: &model.KcalUnit{
							Gr100:    fullIngredient.Kcal.Gr100,
							UnitGr:   fullIngredient.Kcal.UnitGr,
							UnitName: fullIngredient.Kcal.UnitName,
						},
					},
				})
			}
			return ingredientsMapped
		}

		for stgName, ingredients := range stagesDto {
			stages[stgName] = mapIngredients(ingredients)
		}

		return stages
	}

	mapSteps := func(stagesDto map[model.StageName][]*step) map[string][]*model.Step {
		resultSteps := make(map[string][]*model.Step)
		mapStepList := func(steps []*step) []*model.Step {
			stepList := []*model.Step{}
			for _, singleStep := range steps {
				stepList = append(stepList, &model.Step{
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

	recipe := &model.Recipe{
		Id:        &recipeDb.Id,
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
