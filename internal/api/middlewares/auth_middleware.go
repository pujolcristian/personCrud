package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware verifica que el cliente esté autenticado mediante un token.
// Si el token es inválido o está ausente, el middleware detiene la solicitud y devuelve un error 401.
// Este middleware debe aplicarse a rutas protegidas.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token de autorización

		token := c.GetHeader("Authorization")

		// Verificar si el token es válido
		if token != "Bearer 1001-1001-1001" {
			fmt.Print(token)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Detiene la solicitud y no pasa al controlador
			return
		}
		// Continuar al siguiente middleware o controlador
		c.Next()
	}
}
