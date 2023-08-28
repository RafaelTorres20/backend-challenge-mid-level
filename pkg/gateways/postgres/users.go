package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/users"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository struct {
	db *sql.DB
}

// Create implements users.UserRepo.
func (u *UsersRepository) Create(ctx context.Context, user *users.User) error {
	stmt, err := u.db.Prepare("insert into users (id, email, password) values ($1, $2, $3)")
	if err != nil {
		fmt.Println(err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Println(err)
		return err
	}

	id := xid.New().String()

	_, err = stmt.Exec(id, user.Email, hashedPassword)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// DeleteByID implements users.UserRepo.
func (u *UsersRepository) DeleteByID(ctx context.Context, id string) error {
	stmt, err := u.db.Prepare("delete from users where id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}

	tx, err := u.db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = tx.Stmt(stmt).Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

// GetByEmail implements users.UserRepo.
func (u *UsersRepository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	row := u.db.QueryRow("select * from users where email = $1", email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	user := new(users.User)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil

}

// GetByID implements users.UserRepo.
func (u *UsersRepository) GetByID(ctx context.Context, id string) (*users.User, error) {
	row := u.db.QueryRow("select * from users where id = $1", id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	user := new(users.User)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil

}

// UpdateByID implements users.UserRepo.
func (u *UsersRepository) UpdateByID(ctx context.Context, id string, user *users.User) error {
	stmt, err := u.db.Prepare("update users set email = $1, password = $2 where id = $3")
	if err != nil {
		return err
	}

	tx, err := u.db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = tx.Stmt(stmt).Exec(user.Email, hashedPassword, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func NewUsersRepository(db *sql.DB) users.UserRepo {
	return &UsersRepository{db: db}
}
