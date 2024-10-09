package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"udemy_slack_app/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (string, error)
	Read(ctx context.Context, userID string) (*model.User, error)
	Update(ctx context.Context, user *model.User, userID string) error
	Delete(ctx context.Context, userID string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(ctx context.Context, user *model.User) (string, error) {
	result, err := ur.db.Exec("INSERT INTO user (name, age, email, created_at) VALUES (?, ?, ?, ?)", user.Name, user.Age, user.Email, user.CreatedAt)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	idStr := strconv.FormatInt(id, 10)
	return idStr, nil
}

func (ur *userRepository) Read(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := ur.db.QueryRow("SELECT name, age, email FROM user WHERE id = ?", userID).Scan(&user.Name, &user.Age, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Update(ctx context.Context, user *model.User, userID string) error {
	result, err := ur.db.Exec("UPDATE user SET name = ?, age = ?, email = ?, updated_at = ? WHERE id = ?", user.Name, user.Age, user.Email, user.UpdatedAt, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", userID)
	}

	return nil
}

func (ur *userRepository) Delete(ctx context.Context, userID string) error {
	result, err := ur.db.Exec("DELETE FROM user WHERE id = ?", userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", userID)
	}

	return nil
}
