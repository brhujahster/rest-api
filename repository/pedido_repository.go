package repository

import (
	"database/sql"
	"rest-api2/entity"
)

type PedidoRepository struct {
	db *sql.DB
}

func NewPedidoRepository() *PedidoRepository {
	return &PedidoRepository{
		db: nil,
	}
}

func (p *PedidoRepository) GetPedidos() ([]entity.Pedido, error) {
	return []entity.Pedido{}, nil
}
