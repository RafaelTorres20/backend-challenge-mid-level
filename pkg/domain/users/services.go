package users

import (
	"context"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string = os.Getenv("JWT_SECRET")

type service struct {
	userRepo UserRepo
}

// GenerateJWT implements UserService.
func (*service) GenerateJWT(email string) (string, error) {

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
		Subject:   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

// Login implements UserService.
func (u *service) Login(ctx context.Context, email string, password string) (*User, string, error) {
	if email == "" {
		return nil, "", ErrInvalidUser
	}

	if password == "" {
		return nil, "", ErrInvalidPassword
	}

	user, err := u.GetByEmail(ctx, email)

	if user == nil {
		return nil, "", ErrUserNotFound
	}

	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", err
	}

	token, err := u.GenerateJWT(user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
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

func NewUsersService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}
