package users

import (
	model "fit.synapse/przepisnik/app/model/dto"
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
	storage UsersStorage
}

func InitializeUsersModule(basePath string) Facade {
	return &usersFacade{
		storage: storage.NewUsersPersistentStorage(basePath),
	}
}

func (uf *usersFacade) Create(name string, pic string) (*uuid.UUID, error) {
	return uf.storage.Create(name, pic)
}

func (uf *usersFacade) Get(id uuid.UUID) *model.User {
	return uf.storage.Get(id)
}

func (uf *usersFacade) GetAll() []*model.User {
	return uf.storage.GetAll()
}

func (uf *usersFacade) Update(id uuid.UUID, name string, pic string) (*model.User, error) {
	return uf.storage.Update(id, name, pic)
}

func (uf *usersFacade) Delete(id uuid.UUID) error {
	return uf.storage.Delete(id)
}
