package recipes

import (
	model "fit.synapse/przepisnik/app/model/dto"
	"github.com/google/uuid"
)

type RecipesStorage interface {
	Upsert(*model.Recipe) (*uuid.UUID, error)
	Get(id uuid.UUID) *model.Recipe
	// GetAll() []*model.User
	// Update(id uuid.UUID, name string, pic string) (*model.User, error)
	// Delete(id uuid.UUID) error
}
