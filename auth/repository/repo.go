package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/gRPC-example/common"
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
	GetUserByID(c context.Context, id uint) (User, error)
	GetUserRoles(c context.Context, id uint) ([]int, error)
	CreateUser(c context.Context, user User) error
}

func (r *UserRepository) GetUserByEmail(c context.Context, email string) (User, error) {
	user := User{}
	query := `select id , email , password from users where email = $1`
	if err := r.db.QueryRow(c, query, email).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(c context.Context, id uint) (User, error) {
	user := User{}
	query := `select id , email , password from users where id = $1`
	if err := r.db.QueryRow(c, query, id).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetUserRoles(c context.Context, id uint) ([]int, error) {
	roles := []int{}
	query := `select role_id from user_roles where user_id = $1`
	rows, err := r.db.Query(c, query, id)
	if err != nil {
		return roles, err
	}
	for rows.Next() {
		role := 0
		if err := rows.Scan(&role); err != nil {
			return roles, err
		}
		roles = append(roles, role)

	}
	return roles, nil
}

func (r *UserRepository) CreateUser(c context.Context, user User) error {
	var id uint
	query := `INSERT INTO users (email, password) VALUES ($1, $2 ) returning id;`
	if err := r.db.QueryRow(c, query, user.Email, user.Password).Scan(&id); err != nil {
		return err
	}
	query = `insert into user_roles values($1,$2);`
	if _, err := r.db.Exec(c, query, id, common.USER); err != nil {
		return err
	}
	return nil
}
