package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"udemy_slack_app/model"
)

type MessageRepository interface {
	Create(ctx context.Context, message *model.Message) (string, error)
	ReadAll(ctx context.Context) ([]model.Message, error)
	Update(ctx context.Context, message *model.Message, messageID string) error
	Delete(ctx context.Context, messageID string) error
	DeleteAll(ctx context.Context, channelID string) error
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) Create(ctx context.Context, message *model.Message) (string, error) {
	result, err := r.db.Exec("INSERT INTO message (channel_id, user_id, message, created_at) VALUES (?, ?, ?. ?)", message.ChannelID, message.UserID, message.Message, message.CreatedAt)
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

func (r *messageRepository) ReadAll(ctx context.Context) ([]model.Message, error) {
	messages := []model.Message{}
	rows, err := r.db.Query("SELECT * FROM message ORDER BY id ASC")
	if err != nil {
		return messages, err
	}

	for rows.Next() {
		var message model.Message
		if err := rows.Scan(&message.ChannelID, &message.UserID, &message.Message, &message.CreatedAt); err != nil {
			return messages, nil
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func (r *messageRepository) Update(ctx context.Context, message *model.Message, messageID string) error {
	result, err := r.db.Exec("UPDATE message SET message = ?, updated_at = ? WHERE id = ?", message.Message, message.UpdatedAt, messageID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", messageID)
	}

	return nil
}

func (r *messageRepository) Delete(ctx context.Context, messageID string) error {
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	result, err := db.Exec("DELETE FROM message WHERE id = ?", messageID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", messageID)
	}

	return nil
}

func (r *messageRepository) DeleteAll(ctx context.Context, channelID string) error {
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	result, err := db.Exec("DELETE FROM message WHERE channel_id = ?", channelID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", channelID)
	}

	return nil
}
