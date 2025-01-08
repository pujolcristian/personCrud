package controllers

import (
	"net/http"
	"personCrud/internal/domain/dtos"
	"personCrud/internal/domain/models"
	"personCrud/internal/domain/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	useCase *usecases.PersonUseCase
}

// NewPersonController crea una nueva instancia del controlador de Persona.
// Se utiliza para manejar las solicitudes relacionadas con la entidad Persona.
// Parámetros:
//   - useCase: El caso de uso que contiene la lógica de negocio para las Personas.
//
// Retorna:
//   - Un puntero a una instancia de PersonController.
func NewPersonController(useCase *usecases.PersonUseCase) *PersonController {
	return &PersonController{useCase}
}

// Create godoc
// @Summary Crea una nueva persona
// @Description Crea una nueva persona en el sistema
// @Tags Personas
// @Accept json
// @Produce json
// @Param person body models.Person true "Detalles de la persona"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /persons [post]
func (pc *PersonController) Create(c *gin.Context) {
	var body models.Person
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("Invalid request format", err.Error()))
		return
	}

	if err := pc.useCase.CreatePerson(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("Failed to create person", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse("Person created successfully", nil))
}

// GetAll godoc
// @Summary Obtiene todas las personas
// @Description Devuelve una lista de todas las personas registradas
// @Tags Personas
// @Produce json
// @Success 200 {array} models.Person
// @Failure 500 {object} map[string]string
// @Router /persons [get]
func (pc *PersonController) GetAll(c *gin.Context) {
	persons, err := pc.useCase.GetAllPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse("Failed to fetch persons", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse("Persons fetched successfully", persons))
}

// GetByID godoc
// @Summary Obtiene una persona por ID
// @Description Devuelve los detalles de una persona específica
// @Tags Personas
// @Produce json
// @Param id path int true "ID de la persona"
// @Success 200 {object} models.Person
// @Failure 404 {object} map[string]string
// @Router /persons/{id} [get]
func (pc *PersonController) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("Invalid ID", err.Error()))
		return
	}

	person, err := pc.useCase.GetPersonByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dtos.ErrorResponse("Person not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse("Person fetched successfully", person))
}

// Update godoc
// @Summary Actualiza una persona existente
// @Description Actualiza los detalles de una persona específica por ID
// @Tags Personas
// @Accept json
// @Produce json
// @Param id path int true "ID de la persona"
// @Param person body models.Person true "Detalles actualizados de la persona"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid ID or request body"
// @Failure 500 {object} map[string]string "Failed to update person"
// @Router /persons/{id} [put]
func (pc *PersonController) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("Invalid ID", err.Error()))
		return
	}

	var body models.Person
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("Invalid request body", err.Error()))
		return
	}

	err = pc.useCase.UpdatePerson(int32(id), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse("Failed to update person", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse("Person updated successfully", nil))
}

// Delete godoc
// @Summary Elimina una persona
// @Description Elimina los detalles de una persona específica por ID
// @Tags Personas
// @Param id path int true "ID de la persona"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 500 {object} map[string]string "Failed to delete person"
// @Router /persons/{id} [delete]
func (pc *PersonController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse("Invalid ID", err.Error()))
		return
	}

	err = pc.useCase.DeletePerson(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse("Failed to delete person", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse("Person deleted successfully", nil))
}
