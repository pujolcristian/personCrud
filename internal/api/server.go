package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// NewRouter TIP <p> Crea una instancia del router de Gin</p>
func NewRouter() *gin.Engine {
	return gin.Default()
}

// StartServer TIP <p> Arranca el servidor Gin</p>
func StartServer(router *gin.Engine) {
	err := router.Run(":8080")
	if err != nil {
		fmt.Print(err)
		return
	}
}
