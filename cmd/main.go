package main

import (
	"rest-api2/router"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	router.Init(server)
}
