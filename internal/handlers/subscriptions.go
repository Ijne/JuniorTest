package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ijne/JuniorTest/internal/models"
	"github.com/Ijne/JuniorTest/internal/storage/postgres"
)

func CreateSubHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var subscription models.Subscription
			if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
				log.Printf("ERROR FROM[CreateSubHandler] json decode err: %s\n", err)
				return
			}

			if _, err := s.SubscriptionsRepo.Create(subscription.Service_name, subscription.Price, subscription.User_id, subscription.Start_date, subscription.End_date); err != nil {
				log.Printf("ERROR FROM[CreateSubHandler] creating subscription err: %s\n", err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Created"))
		default:
			log.Printf("HTTP: not allowed method %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	}
}

func GetSubHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			id := r.URL.Query().Get("id")

			subscription, err := s.SubscriptionsRepo.Get(id)
			if err != nil {
				log.Printf("ERROR FROM[GetSubHandler] geting subscription err: %s\n", err)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			if err := json.NewEncoder(w).Encode(subscription); err != nil {
				log.Printf("ERROR FROM[GetSubHandler] json decode err: %s\n", err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		default:
			log.Printf("HTTP: not allowed method %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	}
}

func UpdateSubHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			id := r.URL.Query().Get("id")

			var subscription models.Subscription
			if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
				log.Printf("ERROR FROM[UpdateSubHandler] json decode err: %s\n", err)
				return
			}

			if _, err := s.SubscriptionsRepo.Update(id, subscription.Service_name, subscription.Price, subscription.User_id, subscription.Start_date, subscription.End_date); err != nil {
				log.Printf("ERROR FROM[UpdateSubHandler] updating subscription err: %s\n", err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		default:
			log.Printf("HTTP: not allowed method %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	}
}

func DeleteSubHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			id := r.URL.Query().Get("id")

			if _, err := s.SubscriptionsRepo.Delete(id); err != nil {
				log.Printf("ERROR FROM[DeleteSubHandler] deleting subscription err: %s\n", err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		default:
			log.Printf("HTTP: not allowed method %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	}
}

func ListSubHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			subscriptions, err := s.SubscriptionsRepo.GetAll()
			if err != nil {
				log.Printf("ERROR FROM[CreateSubHandler] geting list of subscriptions err: %s\n", err)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			if err := json.NewEncoder(w).Encode(subscriptions); err != nil {
				log.Printf("ERROR FROM[GetSubHandler] json encode err: %s\n", err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		default:
			log.Printf("HTTP: not allowed method %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	}
}

func AmountSubHandler(s *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			service_name := r.URL.Query().Get("service_name")
			user_id := r.URL.Query().Get("user_id")

			amount, err := s.SubscriptionsRepo.GetAmount(service_name, user_id)
			if err != nil {
				log.Printf("ERROR FROM[AmountSubHandler] geting amount of subscriptions err: %s\n", err)
				return
			}

			var data = struct {
				Amount int64
			}{
				Amount: amount,
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			if err := json.NewEncoder(w).Encode(data); err != nil {
				log.Printf("ERROR FROM[AmountSubHandler] json encode err: %s\n", err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Success"))
		default:
			log.Printf("HTTP: not allowed method %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	}
}
