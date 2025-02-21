package controller

import (
	"net/http"
	"rest-api2/entity"
	"rest-api2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	ExempleUseCase *service.ExempleService
}

func NewExampleController(usecase *service.ExempleService) *ExampleController {
	return &ExampleController{
		ExempleUseCase: usecase,
	}
}

func (p *ExampleController) InitRoutes(r *gin.RouterGroup) {
	controller := r.Group("/examples")
	controller.GET("/", p.getExemplos)
	controller.GET("/:id", p.getById)
	controller.POST("/", p.create)
	controller.PUT("/:id", p.update)
	controller.DELETE("/:id", p.delete)
}

func (p *ExampleController) getExemplos(ctx *gin.Context) {
	Exemples, err := p.ExempleUseCase.GetExemples()
	if err != nil {
		p.sendErrorMessage(http.StatusInternalServerError, err.Error(), ctx)
	}
	ctx.JSON(http.StatusOK, Exemples)
}

func (p *ExampleController) getById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		p.sendErrorMessage(http.StatusBadRequest, "Id inválido", ctx)
	}
	example1, err := p.ExempleUseCase.GetExemple(idInt)
	if err != nil {
		p.sendErrorMessage(http.StatusNotFound, "Exemplo não encontrado", ctx)
	}
	ctx.JSON(http.StatusOK, example1)
}

func (p *ExampleController) create(ctx *gin.Context) {
	var Exemplo entity.Exemple
	err := ctx.BindJSON(&Exemplo)
	if err != nil {
		p.sendErrorMessage(http.StatusBadRequest, "Erro ao tentar criar Exemplo", ctx)
	}
	ExemploSalvo, err := p.ExempleUseCase.Create(Exemplo)
	if err != nil {
		p.sendErrorMessage(http.StatusInternalServerError, err.Error(), ctx)
	}
	ctx.JSON(http.StatusOK, ExemploSalvo)
}

func (p *ExampleController) update(ctx *gin.Context) {
	var Exemplo entity.Exemple
	id := ctx.Param("id")
	if id == "" {
		p.sendErrorMessage(http.StatusBadRequest, "Id não informado", ctx)
		return
	}
	ExemploId, err := strconv.Atoi(id)
	if err != nil {
		p.sendErrorMessage(http.StatusBadRequest, "Id inválido", ctx)
		return
	}
	err = ctx.BindJSON(&Exemplo)
	if err != nil {
		p.sendErrorMessage(http.StatusBadRequest, "Erro ao tentar criar Exemplo", ctx)
		return
	}

	exemploAtualizado, err := p.ExempleUseCase.Update(ExemploId, Exemplo)
	if err != nil {
		p.sendErrorMessage(http.StatusInternalServerError, err.Error(), ctx)
	}
	ctx.JSON(http.StatusOK, exemploAtualizado)
}

func (p *ExampleController) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		p.sendErrorMessage(http.StatusBadRequest, "Id não informado", ctx)
		return
	}
	exampleId, err := strconv.Atoi(id)
	if err != nil {
		p.sendErrorMessage(http.StatusBadRequest, "Id inválido", ctx)
		return
	}

	err = p.ExempleUseCase.Delete(exampleId)
	if err != nil {
		p.sendErrorMessage(http.StatusInternalServerError, err.Error(), ctx)
		return
	}
	response := entity.ResponseMessage{
		Message: "Pedido deletado com sucesso",
	}
	ctx.JSON(http.StatusOK, response)
}

func (e ExampleController) sendErrorMessage(httpStatus int, message string, ctx *gin.Context) {
	response := entity.ResponseMessage{
		Message: message,
	}
	ctx.JSON(httpStatus, response)
}
