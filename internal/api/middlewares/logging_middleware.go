package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	// Crear o abrir el archivo de logs
	file, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error al abrir o crear el archivo de logs: %v", err)
	}

	// Configurar el logger para escribir en el archivo
	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		startTime := time.Now()

		// Continuar con el siguiente middleware o controlador
		c.Next()

		// Registrar los detalles de la solicitud
		logger.Printf("[%s] %s %s %s %d %s",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Header,
			c.Writer.Status(),
			time.Since(startTime),
		)
	}
}
