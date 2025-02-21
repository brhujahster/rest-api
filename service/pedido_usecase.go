package service

import (
	"rest-api2/entity"
	model "rest-api2/entity"
	"rest-api2/repository"
)

type PedidoUseCase struct {
	peeidoRepósitory repository.PedidoRepository
}

func NewPedidoUseCase(repository *repository.PedidoRepository) *PedidoUseCase {
	return &PedidoUseCase{}
}

func (p *PedidoUseCase) GetPedidos() ([]entity.Pedido, error) {
	return []model.Pedido{}, nil
}

func (p *PedidoUseCase) GetPedido(id int) (entity.Pedido, error) {
	return model.Pedido{}, nil
}

func (p *PedidoUseCase) Create(pedido entity.Pedido) (entity.Pedido, error) {
	return model.Pedido{}, nil
}

func (p *PedidoUseCase) Update(id int, pedido entity.Pedido) (entity.Pedido, error) {
	return model.Pedido{}, nil
}

func (p *PedidoUseCase) Delete(id int) error {
	return nil
}
