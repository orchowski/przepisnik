package users

import (
	model "fit.synapse/przepisnik/app/model/dto"
	users_public "fit.synapse/przepisnik/app/model/users/public"
	storage "fit.synapse/przepisnik/storage/users"
	"github.com/google/uuid"
)

type Facade interface {
	Create(name string, pic string) (*uuid.UUID, error)
	Get(id uuid.UUID) *model.User
	GetAll() []*model.User
	Update(id uuid.UUID, name string, pic string) (*model.User, error)
	Delete(id uuid.UUID) error
}

type usersFacade struct {
	storage users_public.UsersStorage
}

func InitializeUsersModule(basePath string) Facade {
	return &usersFacade{
		storage: storage.NewUsersStorage(basePath),
	}
}

func (uf *usersFacade) Create(name string, pic string) (*uuid.UUID, error) {
	return nil, nil
}
func (uf *usersFacade) Get(uuid.UUID) *model.User {
	return nil
}
func (uf *usersFacade) GetAll() []*model.User {
	return nil
}
func (uf *usersFacade) Update(id uuid.UUID, name string, pic string) (*model.User, error) {
	return nil, nil
}
func (uf *usersFacade) Delete(id uuid.UUID) error {
	return nil
}
