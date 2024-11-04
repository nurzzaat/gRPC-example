package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID       uint
	Email    string
	Password string
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) AuthRepository {
	return &UserRepository{db: db}
}

type AuthRepository interface {
	GetUserByEmail(c context.Context, email string) (User, error)
}

func (r *UserRepository) GetUserByEmail(c context.Context, email string) (User, error) {
	fmt.Println("Gettin' user from db...")
	return User{ID: 1}, nil
}
