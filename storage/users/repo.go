package storage

import (
	"errors"

	model "fit.synapse/przepisnik/app/model/dto"
	"fit.synapse/przepisnik/commons"
	"github.com/google/uuid"
)

type usersRepo struct {
	users map[uuid.UUID]*user
}

func NewUsersPersistentStorage(basePath string) *usersRepo {
	return NewUsersInMemoryStorage()
}

func NewUsersInMemoryStorage() *usersRepo {
	return &usersRepo{users: make(map[uuid.UUID]*user)}
}

func (uf *usersRepo) Create(name string, pic string) (*uuid.UUID, error) {
	id := uuid.New()
	user := &user{
		id:             id,
		name:           name,
		profilePicPath: pic,
	}
	if validationError := isValid(user); validationError != nil {
		return nil, validationError
	}
	uf.users[id] = user
	return &id, nil
}

func (uf *usersRepo) Get(id uuid.UUID) *model.User {
	user := uf.users[id]
	if user == nil {
		return nil
	}
	return &model.User{
		Id:             id,
		Name:           user.name,
		ProfilePicPath: user.profilePicPath,
	}
}

func (uf *usersRepo) GetAll() []*model.User {
	var allUsers []*model.User
	for _, user := range uf.users {
		allUsers = append(allUsers, &model.User{
			Id:             user.id,
			Name:           user.name,
			ProfilePicPath: user.profilePicPath,
		})
	}
	return allUsers
}

func (uf *usersRepo) Update(id uuid.UUID, name string, pic string) (*model.User, error) {
	user := uf.users[id]
	if user == nil {
		return nil, commons.UserNotFound
	}
	user.name = name
	user.profilePicPath = pic

	return uf.Get(id), nil
}

func (uf *usersRepo) Delete(id uuid.UUID) error {
	user := uf.users[id]
	if user == nil {
		return commons.UserNotFound
	}
	delete(uf.users, id)
	return nil
}

func isValid(user *user) error {
	if len(user.name) == 0 {
		return errors.New("name must be provided")
	}
	return nil
}
