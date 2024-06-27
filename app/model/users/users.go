package users

import (
	model "fit.synapse/przepisnik/app/model/dto"
	"github.com/google/uuid"
)

type UsersStorage interface {
	Create(name string, pic string) (*uuid.UUID, error)
	Get(id uuid.UUID) *model.User
	GetAll() []*model.User
	Update(id uuid.UUID, name string, pic string) (*model.User, error)
	Delete(id uuid.UUID) error
}
