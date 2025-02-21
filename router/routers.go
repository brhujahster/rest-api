package router

import (
	"rest-api2/controller"
	"rest-api2/repository"
	"rest-api2/service"

	"github.com/gin-gonic/gin"
)

// initializeRoutes sets up the routes for the Pedido resource in the Gin router.
// It creates a new PedidoRepository, PedidoUseCase, and PedidoController, and
// registers the routes for the Pedido resource under the "/pedidos" path.
func InitializeRoutes(router *gin.RouterGroup) {
	exempleRepository := repository.NewPedidoRepository()
	examapleService := service.NewExempleService(exempleRepository)
	pedidoController := controller.NewExampleController(examapleService)

	pedidoController.InitRoutes(router)

}
