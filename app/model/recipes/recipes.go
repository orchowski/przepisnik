package recipes

import (
	"fmt"

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
    AddTags(recipeId uuid.UUID, tags []string) (*model.Recipe, error)
    SetCalories(recipeId uuid.UUID, calories int) (*model.Recipe, error)
    CategorizeByCalories(recipes []model.Recipe) (map[string][]model.Recipe, error)
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

func (r *recipes) AddTags(recipeId uuid.UUID, tags []string) (*model.Recipe, error) {
    return nil, nil
}

func (r *recipes) SetCalories(recipeId uuid.UUID, calories int) (*model.Recipe, error) {
    return nil, nil
}

func (r *recipes) CategorizeByCalories(recipes []model.Recipe) (map[string][]model.Recipe, error) {
    categories := map[string][]model.Recipe{
        "<300":       {},
        "300-499":    {},
        "500-749":    {},
        "750-999":    {},
        "1000-1500":  {},
        ">1500":      {},
    }

    for _, recipe := range recipes {
        if recipe.Calories < 0 {
            return nil, fmt.Errorf("invalid calorie value: %d", recipe.Calories)
        }
        switch {
        case recipe.Calories < 300:
            categories["<300"] = append(categories["<300"], recipe)
        case recipe.Calories >= 300 && recipe.Calories <= 499:
            categories["300-499"] = append(categories["300-499"], recipe)
        case recipe.Calories >= 500 && recipe.Calories <= 749:
            categories["500-749"] = append(categories["500-749"], recipe)
        case recipe.Calories >= 750 && recipe.Calories <= 999:
            categories["750-999"] = append(categories["750-999"], recipe)
        case recipe.Calories >= 1000 && recipe.Calories <= 1500:
            categories["1000-1500"] = append(categories["1000-1500"], recipe)
        case recipe.Calories > 1500:
            categories[">1500"] = append(categories[">1500"], recipe)
        }
    }

    return categories, nil
}