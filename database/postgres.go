package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"platzi.com/go/cqrs/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

func (repo *PostgresRepository) InsertFeed(ctx context.Context, feed *models.Feed) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO feeds (id, title, description) VALUES ($1, $2, $3)", feed.ID, feed.Title, feed.Description)
	return err
}

func (repo *PostgresRepository) ListFeeds(ctx context.Context) ([]*models.Feed, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, title, description, created_at FROM feeds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := make([]*models.Feed, 0)
	for rows.Next() {
		var id, title, description string
		var createdAt time.Time
		if err := rows.Scan(&id, &title, &description, &createdAt); err != nil {
			return nil, err
		}

		feeds = append(feeds, &models.Feed{
			ID:          id,
			Title:       title,
			Description: description,
			CreatedAt:   createdAt,
		})
	}

	return feeds, nil
}
