package controller

import (
	"net/http"
	"rest-api2/entity"
	"rest-api2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pedidoController struct {
	pedidoUseCase *service.PedidoUseCase
}

func NewPedidoController(usecase *service.PedidoUseCase) *pedidoController {
	return &pedidoController{
		pedidoUseCase: usecase,
	}
}

func (p *pedidoController) InitRoutes(r *gin.RouterGroup) {
	controller := r.Group("/pedidos")
	controller.GET("/", p.getPedidos)
	controller.GET("/:id", p.getById)
	controller.POST("/", p.create)
	controller.PUT("/:id", p.update)
	controller.DELETE("/:id", p.delete)
}

func (p *pedidoController) getPedidos(ctx *gin.Context) {
	pedidos, err := p.pedidoUseCase.GetPedidos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, pedidos)
}

func (p *pedidoController) getById(ctx *gin.Context) {
	var pedido entity.Pedido
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id inválido",
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}
	pedido, err = p.pedidoUseCase.GetPedido(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, pedido)
}

func (p *pedidoController) create(ctx *gin.Context) {
	var pedido entity.Pedido
	err := ctx.BindJSON(&pedido)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	pedidoSalvo, err := p.pedidoUseCase.Create(pedido)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, pedidoSalvo)
}

func (p *pedidoController) update(ctx *gin.Context) {
	var pedido entity.Pedido
	id := ctx.Param("id")
	if id == "" {
		response := entity.ResponseMessage{
			Message: "Id não informado",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pedidoId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id inválido",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = ctx.BindJSON(&pedido)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id não informado",
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	pedidoAtualizado, err := p.pedidoUseCase.Update(pedidoId, pedido)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, pedidoAtualizado)
}

func (p *pedidoController) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := entity.ResponseMessage{
			Message: "Id não informado",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pedidoId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id inválido",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.pedidoUseCase.Delete(pedidoId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	response := entity.ResponseMessage{
		Message: "Pedido deletado com sucesso",
	}
	ctx.JSON(http.StatusOK, response)
}
