package service

import (
	"rest-api2/entity"
	"rest-api2/repository"
)

type ExempleService struct {
	pedidoRepository repository.ExampleRepository
}

func NewExempleService(repository *repository.ExampleRepository) *ExempleService {
	return &ExempleService{
		pedidoRepository: *repository,
	}
}

func (p *ExempleService) GetExemples() ([]entity.Exemple, error) {
	return []entity.Exemple{}, nil
}

func (p *ExempleService) GetExemple(id int) (entity.Exemple, error) {
	return entity.Exemple{}, nil
}

func (p *ExempleService) Create(exemple entity.Exemple) (entity.Exemple, error) {
	return entity.Exemple{}, nil
}

func (p *ExempleService) Update(id int, Exemplo entity.Exemple) (entity.Exemple, error) {
	return entity.Exemple{}, nil
}

func (p *ExempleService) Delete(id int) error {
	return nil
}
