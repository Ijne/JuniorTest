package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Ijne/JuniorTest/config"
	"github.com/Ijne/JuniorTest/internal/storage/repos"
	_ "github.com/lib/pq"
)

type Storage struct {
	db                *sql.DB
	SubscriptionsRepo repos.SubscriptionsRepo
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	db, err := sql.Open("postgres", getDSN(cfg))
	if err != nil {
		return nil, err
	}

	subscriptionsRepo, err := NewSubscriptionsRepo(db)
	if err != nil {
		return nil, fmt.Errorf("starting subscriptions repo: %w", err)
	}

	return &Storage{db: db, SubscriptionsRepo: subscriptionsRepo}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func getDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
		cfg.Postgres.SSLMode,
	)
}
