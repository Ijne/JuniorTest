package repos

import "github.com/Ijne/JuniorTest/internal/models"

type SubscriptionsRepo interface {
	Create(service_name string, price int64, user_id string, start_date, end_date string) (int64, error)
	Get(id string) (models.Subscription, error)
	Update(id string, service_name string, price int64, user_id string, start_date, end_date string) (string, error)
	Delete(id string) (string, error)
	GetAll() (*[]models.Subscription, error)
}
