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
	Stages    map[string][]Ingredient // stage is ex. sauce
	Steps     []Step
	PhotoURLs []string
}

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
	Kcal  *Kcal
}

type Kcal struct {
	Value int
	Unit  KcalUnit
}

type KcalUnit struct {
	Gr100    int
	UnitGr   *int
	UnitName *string
}

type Comment struct {
	User    User
	Date    time.Time
	Content string
}
