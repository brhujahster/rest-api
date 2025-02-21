package controller

import (
	"net/http"
	"rest-api2/entity"
	"rest-api2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExempleController struct {
	ExempleUseCase *service.ExempleService
}

func NewExampleController(usecase *service.ExempleService) *ExempleController {
	return &ExempleController{
		ExempleUseCase: usecase,
	}
}

func (p *ExempleController) InitRoutes(r *gin.RouterGroup) {
	controller := r.Group("/exemples")
	controller.GET("/", p.getExemplos)
	controller.GET("/:id", p.getById)
	controller.POST("/", p.create)
	controller.PUT("/:id", p.update)
	controller.DELETE("/:id", p.delete)
}

func (p *ExempleController) getExemplos(ctx *gin.Context) {
	Exemples, err := p.ExempleUseCase.GetExemples()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, Exemples)
}

func (p *ExempleController) getById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id inválido",
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}
	example1, err := p.ExempleUseCase.GetExemple(idInt)

	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id inválido",
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	ctx.JSON(http.StatusOK, example1)
}

func (p *ExempleController) create(ctx *gin.Context) {
	var Exemplo entity.Exemple
	err := ctx.BindJSON(&Exemplo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ExemploSalvo, err := p.ExempleUseCase.Create(Exemplo)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, ExemploSalvo)
}

func (p *ExempleController) update(ctx *gin.Context) {
	var Exemplo entity.Exemple
	id := ctx.Param("id")
	if id == "" {
		response := entity.ResponseMessage{
			Message: "Id não informado",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ExemploId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id inválido",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = ctx.BindJSON(&Exemplo)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id não informado",
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	exemploAtualizado, err := p.ExempleUseCase.Update(ExemploId, Exemplo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, exemploAtualizado)
}

func (p *ExempleController) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := entity.ResponseMessage{
			Message: "Id não informado",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	exampleId, err := strconv.Atoi(id)
	if err != nil {
		response := entity.ResponseMessage{
			Message: "Id inválido",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.ExempleUseCase.Delete(exampleId)

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
