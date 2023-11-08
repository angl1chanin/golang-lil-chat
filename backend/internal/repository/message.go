package repository

import (
	"chat/internal/entity"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type MessageRepository interface {
	Create(author, text string) error
	Get(limit int) ([]entity.Message, error)
	Delete(id int) error
}

type messageRepository struct {
	db *sql.DB
}

var _ MessageRepository = (*messageRepository)(nil)

func New(storagePath string) (*messageRepository, error) {
	const op = "chat.internal.repository.message_repository.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS messages(
		    id INTEGER PRIMARY KEY,
		    author VARCHAR(25) NOT NULL,
		    text TEXT NOT NULL,
		    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &messageRepository{db: db}, nil
}

func (r *messageRepository) Create(author, text string) error {
	const op = "chat.internal.repository.message_repository.Create"

	stmt, err := r.db.Prepare(`INSERT INTO messages(author, text, timestamp) VALUES(?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(author, text, time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *messageRepository) Get(limit int) ([]entity.Message, error) {
	const op = "chat.internal.repository.message_repository.Get"

	var messages []entity.Message

	rows, err := r.db.Query(`SELECT * FROM messages ORDER BY id DESC LIMIT ? OFFSET 0`, limit)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()
	for rows.Next() {
		var message entity.Message
		err = rows.Scan(&message.ID, &message.Author, &message.Text, &message.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (r *messageRepository) Delete(id int) error {
	const op = "chat.internal.repository.message_repository.Delete"

	stmt, err := r.db.Prepare(`DELETE FROM messages WHERE id=?`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
