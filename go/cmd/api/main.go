package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Order struct {
	Event   string `json:"event"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

type OrderService struct {
	repo *OrderRepository
}

type OrderRepository struct {
	db *sql.DB
}

func main() {

}

func GetOrderHandler(svc *OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.PathValue("id")
		order, err := svc.GetOrder(ctx, id)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, err)
			return
		}
		writeJSON(w, http.StatusOK, order)
	}
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*Order, error) {
	return s.repo.GetOrder(ctx, id)
}

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*Order, error) {
	const query = `
		select event, id, message
		from orders
		where id = $1
	`
	var order Order

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&order.Event,
		&order.Id,
		&order.Message,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
