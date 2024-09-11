package storage

import (
	"time"

	"github.com/google/uuid"
)

type recipeId = uuid.UUID
type recipe struct {
	Id        recipeId
	Name      string
	Stages    map[string][]*recipeIngredient
	Steps     map[string][]*step
	PhotoURLs []string
	Author    uuid.UUID
	Portions  int
	Tags      []string
	Calories  int
	Rating    int
	Deleted   bool
}

type step struct {
	Description string
	PhotoURLs   []string
}

type ingredientId = uuid.UUID

type recipeIngredient struct {
	Id          ingredientId
	AmountGr    int
	AmountUnits *float64
}

type ingredient struct {
	Id   ingredientId
	Name string
	Kcal *kcalUnit
}

type kcalUnit struct {
	Gr100    int
	UnitGr   *int
	UnitName *string
}

type comment struct {
}

type recipeRate struct {
	To     uuid.UUID
	User   uuid.UUID
	Date   time.Time
	Rating int
}
