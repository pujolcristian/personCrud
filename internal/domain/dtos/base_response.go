package dtos

// BaseResponse representa la estructura estándar para las respuestas de la API.
// Contiene campos para indicar el éxito, mensajes, datos y errores.
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse genera una respuesta exitosa
func SuccessResponse(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse genera una respuesta de error
func ErrorResponse(message string, err string) BaseResponse {
	return BaseResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
