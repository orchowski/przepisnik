package storage

import "github.com/google/uuid"

type user struct {
	id             uuid.UUID
	name           string
	profilePicPath string
}
