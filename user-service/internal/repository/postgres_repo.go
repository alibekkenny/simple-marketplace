package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alibekkenny/simple-marketplace/user-service/internal/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// GetUserByID(id string) (*model.User, error)
func (r *PostgresRepository) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	stmt := `SELECT id, username, email, role FROM users`

	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser(user *model.User) (int, error)
func (r *PostgresRepository) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	var id int64
	stmt := `INSERT INTO users (username, email, password, role)
	VALUES($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRowContext(ctx, stmt, user.Username, user.Email, user.Password, user.Role).Scan(&id)
	if err != nil {
		if isUniqueViolation(err) {
			return 0, model.ErrDuplicate
		}
		return 0, err
	}

	return id, nil
}

// FindUserByEmail(email string) (*model.User, error)
func (r *PostgresRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	stmt := `SELECT id, username, email, password, role FROM users WHERE email = $1`

	err := r.db.QueryRowContext(ctx, stmt, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

// ExistsByEmailOrUsername(email, username string) (bool, error)
func (r *PostgresRepository) ExistsByEmailOrUsername(ctx context.Context, email, username string) (bool, error) {
	var exists bool
	stmt := `SELECT EXISTS (
		SELECT 1 FROM users WHERE email = $1 OR username = $2
	)`

	err := r.db.QueryRowContext(ctx, stmt, email, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
