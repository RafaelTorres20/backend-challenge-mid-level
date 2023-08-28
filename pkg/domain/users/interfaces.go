package users

import "context"

type Service interface {
	// GetByID returns user by id
	GetByID(ctx context.Context, id string) (*User, error)
	// GetByEmail returns user by email
	GetByEmail(ctx context.Context, email string) (*User, error)
	// Create creates new user
	Create(ctx context.Context, user *User) error
	// Update updates user
	UpdateByID(ctx context.Context, id string, user *User) error
	// Delete deletes user
	DeleteByID(ctx context.Context, id string) error
	// Login logs user in
	Login(ctx context.Context, email, password string) (*User, string, error)
	// GenerateJWT generates JWT token
	GenerateJWT(email string) (string, error)
}

type UserRepo interface {
	// GetByID returns user by id
	GetByID(ctx context.Context, id string) (*User, error)
	// GetByEmail returns user by email
	GetByEmail(ctx context.Context, email string) (*User, error)
	// Create creates new user
	Create(ctx context.Context, user *User) error
	// Update updates user
	UpdateByID(ctx context.Context, id string, user *User) error
	// Delete deletes user
	DeleteByID(ctx context.Context, id string) error
}
