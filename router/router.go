package router

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	v1 := router.Group("/api/v1")

	InitializeRoutes(v1)

	router.Run(":8080")

}
