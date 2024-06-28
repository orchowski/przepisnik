package model

import uuid "github.com/google/uuid"
import "time"

type User struct {
	Id             uuid.UUID
	Name           string
	ProfilePicPath string
}

type Recipe struct {
	Id        uuid.UUID
	Name      string
	Stages    map[string][]Ingredient // stage is ex. sauce and dish itself. Or one stage is to prepare pesto to make a sandwitch, another one bread, finally we can place ingredients there
	Steps     map[string][]Step       // where key is stage
	PhotoURLs []string
	Author    uuid.UUID
	Portions  int
	Tags      []string
}

// todo : search np kategoria i kaloryczność całego dania / porcja / cały posiłek
// obiad, porcja ma : >100, >200, >500

type Step struct {
	Description string
	PhotoURLs   []string
}

type Ingredient struct {
	Id     uuid.UUID
	Name   string
	Amount IngredientAmount
}

type IngredientAmount struct {
	Gr    int
	Units *float32
	Kcal  *KcalUnit
}

type KcalUnit struct {
	Gr100    int
	UnitGr   *int
	UnitName *string
}

type Comment struct {
	To      uuid.UUID
	User    User
	Date    time.Time
	Content string
}

type RecipeRate struct {
	To     uuid.UUID
	User   User
	Date   time.Time
	Rating int
}
