package repository

import (
	"database/sql"
	"rest-api2/entity"
)

type ExampleRepository struct {
	db *sql.DB
}

func NewPedidoRepository() *ExampleRepository {
	return &ExampleRepository{
		db: nil,
	}
}

func (p *ExampleRepository) GetPedidos() ([]entity.Exemple, error) {
	return []entity.Exemple{}, nil
}
