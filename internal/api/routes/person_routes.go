package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"personCrud/internal/api/controllers"
)

// RegisterPersonRoutes registra las rutas relacionadas con la entidad Persona.
// Aplica middlewares específicos para proteger las rutas y manejar CORS.
// Parámetros:
//   - router: La instancia de Gin donde se registrarán las rutas.
//   - personController: El controlador que maneja las solicitudes de Persona.
func RegisterPersonRoutes(
	router *gin.Engine,
	personController *controllers.PersonController,
	params struct {
		fx.In
		AuthMiddleware    gin.HandlerFunc `name:"authMiddleware"`
		CORSMiddleware    gin.HandlerFunc `name:"corsMiddleware"`
		LoggingMiddleware gin.HandlerFunc `name:"loggingMiddleware"`
	},
) {
	personGroup := router.Group("/persons", params.LoggingMiddleware)
	personGroup.GET("/", personController.GetAll)
	authenticatedGroup := personGroup.Group("/",
		params.AuthMiddleware,
		params.CORSMiddleware,
	)
	{
		authenticatedGroup.POST("/", personController.Create)
		authenticatedGroup.GET("/:id", personController.GetByID)
		authenticatedGroup.PUT("/:id", personController.Update)
		authenticatedGroup.DELETE("/:id", personController.Delete)
	}
}
