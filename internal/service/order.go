package service

import (
	"nunu-project/internal/model"
	"nunu-project/internal/repository"
)

type OrderService interface {
	GetOrderById(id int64) (*model.Order, error)
}

type orderService struct {
	*Service
	orderRepository repository.OrderRepository
}

func NewOrderService(service *Service, orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		Service:        service,
		orderRepository: orderRepository,
	}
}

func (s *orderService) GetOrderById(id int64) (*model.Order, error) {
	return s.orderRepository.FirstById(id)
}
