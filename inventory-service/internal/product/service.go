package product

import "inventory-service/models"

type Service struct {
	Repo *Repository
}

func (s *Service) CreateProduct(p *models.Product) error {
	return s.Repo.Create(p)
}

func (s *Service) GetProducts() ([]models.Product, error) {
	return s.Repo.GetAll()
}
func (s *Service) UpdateProduct(id uint, updates map[string]interface{}) error {
	return s.Repo.Update(id, updates)
}
func (s *Service) DeleteProduct(id uint) error {
	return s.Repo.Delete(id)
}
func (s *Service) GetProductByID(id uint) (*models.Product, error) {
	return s.Repo.GetByID(id)
}
