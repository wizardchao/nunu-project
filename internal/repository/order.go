package repository

import (
	"nunu-project/internal/model"
)

type OrderRepository interface {
	FirstById(id int64) (*model.Order, error)
}
type orderRepository struct {
	*Repository
}

func NewOrderRepository(repository *Repository) OrderRepository {
	return &orderRepository{
		Repository: repository,
	}
}

func (r *orderRepository) FirstById(id int64) (*model.Order, error) {
	var order model.Order
	// TODO: query db
	return &order, nil
}
