package storage

import (
	"testing"

	"fit.synapse/przepisnik/commons"
	"github.com/google/uuid"
)

func TestNewUsersStorage(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	if storage == nil {
		t.Fatal("Expected non-nil storage")
	}
	if storage.users == nil {
		t.Fatal("Expected non-nil map of users")
	}
}

func TestCreateUser(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	id, err := storage.Create("John Doe", "/path/to/profilepic")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if id == nil {
		t.Fatal("Expected non-nil ID")
	}
	if _, exists := storage.users[*id]; !exists {
		t.Fatalf("Expected user to be in storage")
	}
}

func TestCreateUserInvalid(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	_, err := storage.Create("", "/path/to/profilepic")

	if err == nil {
		t.Fatal("Expected error for invalid user")
	}
}

func TestGetUser(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	id, _ := storage.Create("John Doe", "/path/to/profilepic")

	user := storage.Get(*id)
	if user == nil {
		t.Fatal("Expected non-nil user")
	}
	if user.Name != "John Doe" {
		t.Fatalf("Expected name to be 'John Doe', got %s", user.Name)
	}
	if user.ProfilePicPath != "/path/to/profilepic" {
		t.Fatalf("Expected profile pic path to be '/path/to/profilepic', got %s", user.ProfilePicPath)
	}
}

func TestGetUserNotFound(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	id := uuid.New()

	user := storage.Get(id)
	if user != nil {
		t.Fatal("Expected nil user for non-existent ID")
	}
}

func TestGetAllUsers(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	storage.Create("John Doe", "/path/to/profilepic")
	storage.Create("Jane Doe", "/path/to/anotherprofilepic")

	users := storage.GetAll()
	if len(users) != 2 {
		t.Fatalf("Expected 2 users, got %d", len(users))
	}
}

func TestUpdateUser(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	id, _ := storage.Create("John Doe", "/path/to/profilepic")

	updatedUser, err := storage.Update(*id, "John Smith", "/new/path/to/profilepic")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if updatedUser == nil {
		t.Fatal("Expected non-nil updated user")
	}
	if updatedUser.Name != "John Smith" {
		t.Fatalf("Expected name to be 'John Smith', got %s", updatedUser.Name)
	}
	if updatedUser.ProfilePicPath != "/new/path/to/profilepic" {
		t.Fatalf("Expected profile pic path to be '/new/path/to/profilepic', got %s", updatedUser.ProfilePicPath)
	}
}

func TestUpdateUserNotFound(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	id := uuid.New()

	_, err := storage.Update(id, "John Smith", "/new/path/to/profilepic")
	if err != commons.UserNotFound {
		t.Fatalf("Expected UserNotFound error, got %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	id, _ := storage.Create("John Doe", "/path/to/profilepic")

	err := storage.Delete(*id)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if _, exists := storage.users[*id]; exists {
		t.Fatalf("Expected user to be deleted from storage")
	}
}

func TestDeleteUserNotFound(t *testing.T) {
	storage := NewUsersInMemoryStorage()
	id := uuid.New()

	err := storage.Delete(id)
	if err != commons.UserNotFound {
		t.Fatalf("Expected UserNotFound error, got %v", err)
	}
}
