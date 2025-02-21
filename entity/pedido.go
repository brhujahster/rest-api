package entity

type Pedido struct {
	ID     int     `json:"id"`
	Status string  `json:"status"`
	Valor  float64 `json:"valor"`
}
