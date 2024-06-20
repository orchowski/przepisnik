package storage

import "github.com/google/uuid"
import "fit.synapse/przepisnik/app/model/users/public"

type usersRepo struct {
	users map[uuid.UUID]user
}

func NewUsersStorage(basePath string) users_public.UsersStorage {
	return &usersRepo{}
}
