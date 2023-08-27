package users

import (
	"context"
)

type service struct {
	userRepo UserRepo
}

// Create implements UserRepo.
func (u *service) Create(ctx context.Context, user *User) error {
	if user.Email == "" {
		return ErrInvalidUser
	}

	if user.Password == "" {
		return ErrInvalidPassword
	}

	return u.userRepo.Create(ctx, user)
}

// Delete implements UserRepo.
func (u *service) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	return u.userRepo.DeleteByID(ctx, id)
}

// GetByEmail implements UserRepo.
func (u *service) GetByEmail(ctx context.Context, email string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidUser
	}

	return u.userRepo.GetByEmail(ctx, email)
}

// GetByID implements UserRepo.
func (u *service) GetByID(ctx context.Context, id string) (*User, error) {
	if id == "" {
		return nil, ErrInvalidID
	}

	return u.userRepo.GetByID(ctx, id)
}

// Update implements UserRepo.
func (u *service) UpdateByID(ctx context.Context, id string, user *User) error {
	if user.Email == "" {
		return ErrInvalidUser
	}

	if user.Password == "" {
		return ErrInvalidPassword
	}

	return u.userRepo.UpdateByID(ctx, id, user)
}

func NewUsersService(userRepo UserRepo) UserService {
	return &service{
		userRepo: userRepo,
	}
}
