package order

import (
	"order-service/models"
)

type Service struct {
	Repo *Repository
}

func (s *Service) CreateOrder(order *models.Order) error {
	return s.Repo.Create(order)
}

func (s *Service) GetOrderByID(id uint) (*models.Order, error) {
	return s.Repo.GetByID(id)
}

func (s *Service) ListOrders() ([]models.Order, error) {
	return s.Repo.GetAll()
}

func (s *Service) UpdateOrder(order *models.Order) error {
	updates := map[string]interface{}{
		"user_id":     order.UserID,
		"product_id":  order.ProductID,
		"quantity":    order.Quantity,
		"total_price": order.TotalPrice,
		"status":      order.Status,
	}
	return s.Repo.Update(order.ID, updates)
}

func (s *Service) DeleteOrder(id uint) error {
	return s.Repo.Delete(id)
}
