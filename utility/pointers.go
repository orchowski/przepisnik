package utility

import "github.com/google/uuid"

type NotNil interface {
	~int | ~float32 | ~float64 | ~string | bool
}

func Newuuid() *uuid.UUID {
	newUUID := uuid.New()
	return &newUUID
}

func NewFloat32(v float32) *float32 {
	return &v
}
func NewInt(v int) *int {
	return &v
}

func AddressOf[T NotNil](v T) *T {
	return &v
}
