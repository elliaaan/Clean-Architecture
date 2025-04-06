package order

import "order-service/models"

type Service struct {
	Repo *Repository
}

func (s *Service) CreateOrder(order *models.Order) error {
	return s.Repo.Create(order)
}

func (s *Service) GetOrders() ([]models.Order, error) {
	return s.Repo.GetAll()
}

func (s *Service) GetOrderByID(id uint) (*models.Order, error) {
	return s.Repo.GetByID(id)
}

func (s *Service) UpdateOrderStatus(id uint, status string) error {
	return s.Repo.UpdateStatus(id, status)
}
func (s *Service) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.Repo.GetByUserID(userID)
}
