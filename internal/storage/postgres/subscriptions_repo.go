package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Ijne/JuniorTest/internal/models"
)

type SubscriptionsRepo struct {
	db *sql.DB
}

func NewSubscriptionsRepo(db *sql.DB) (*SubscriptionsRepo, error) {
	repo := SubscriptionsRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *SubscriptionsRepo) CreateTable() error {
	const op = "postgres.subscriptions.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS public.subscriptions (
			id bigserial NOT NULL,
			service_name varchar NOT NULL,
			price numeric NOT NULL,
			user_id uuid NOT NULL,
			start_date varchar NOT NULL,
			end_date varchar NOT NULL,
			CONSTRAINT subscriptions_pk PRIMARY KEY (id)
		);
	`

	_, err := r.db.Exec(stmt)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}
	return err
}

func (r *SubscriptionsRepo) Create(service_name string, price int64, user_id string, start_date, end_date string) (int64, error) {
	const op = "postgres.subscriptions.create"

	var id int64

	err := r.db.QueryRow(`
		INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, service_name, price, user_id, start_date, end_date).Scan(&id)

	if err != nil {
		log.Printf("%s: %v", op, err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, err
}

func (r *SubscriptionsRepo) Get(id string) (models.Subscription, error) {
	const op = "postgres.subscriptions.get"

	var subscription models.Subscription

	err := r.db.QueryRow(`
		SELECT * FROM subscriptions WHERE id = $1
	`, id).Scan(&subscription.ID, &subscription.Service_name, &subscription.Price, &subscription.User_id, &subscription.Start_date, &subscription.End_date)

	if err != nil {
		log.Printf("%s: %v", op, err)
		return models.Subscription{}, fmt.Errorf("%s: %w", op, err)
	}

	return subscription, err
}

func (r *SubscriptionsRepo) Update(id string, service_name string, price int64, user_id string, start_date, end_date string) (string, error) {
	const op = "postgres.subscriptions.update"

	err := r.db.QueryRow(`
		UPDATE subscriptions
		SET service_name = $1, price = $2, start_date = $3, end_date = $4
		WHERE id = $5
		RETURNING id
	`, service_name, price, start_date, end_date, id).Scan(&id)

	if err != nil {
		log.Printf("%s: %v", op, err)
		return id, fmt.Errorf("%s: %w", op, err)
	}

	return id, err
}

func (r *SubscriptionsRepo) Delete(id string) (string, error) {
	const op = "postgres.subscriptions.delete"

	_, err := r.db.Exec(`DELETE FROM subscriptions WHERE id = $1`, id)

	if err != nil {
		log.Printf("%s: %v", op, err)
		return id, fmt.Errorf("%s: %w", op, err)
	}

	return id, err
}

func (r *SubscriptionsRepo) GetAll() (*[]models.Subscription, error) {
	const op = "postgres.subscriptions.create"

	rows, err := r.db.Query(`
		SELECT * FROM subscriptions
	`)

	if err != nil {
		log.Printf("%s: %v", op, err)
		return &[]models.Subscription{}, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		err := rows.Scan(&sub)
		if err != nil {
			log.Printf("%s: %v", op, err)
			return &[]models.Subscription{}, fmt.Errorf("%s: %w", op, err)
		}
		subscriptions = append(subscriptions, sub)
	}

	return &subscriptions, err
}
