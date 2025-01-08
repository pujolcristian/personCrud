package db

import (
	"fmt"
	"gorm.io/gorm"
	"personCrud/internal/domain/models"
	"personCrud/internal/domain/repositories"
)

type GormPersonRepository struct {
	db *gorm.DB
}

func NewGormPersonRepository(db *gorm.DB) repositories.PersonRepository {
	return &GormPersonRepository{db}
}

// Create inserta una nueva Persona en la base de datos.
// Parámetros:
//   - person: Un puntero a la estructura Person que contiene los datos a insertar.
//
// Retorna:
//   - Un error si ocurre un problema durante la inserción en la base de datos.
func (r *GormPersonRepository) Create(person *models.Person) error {
	// Verificar si el usuario ya existe
	var existing models.Person
	if err := r.db.Where("id = ?", person.ID).First(&existing).Error; err == nil {
		return fmt.Errorf("user with ID %d already exists", person.ID)
	}

	// Crear el usuario si no existe
	return r.db.Create(person).Error
}

// GetAll obtiene una lista de todas las Personas registradas en la base de datos.
// Retorna:
//   - Una lista de Personas.
func (r *GormPersonRepository) GetAll() ([]models.Person, error) {
	var persons []models.Person
	err := r.db.Find(&persons).Error
	return persons, err
}

// GetByID obtiene una Persona de la base de datos por su ID.
// Parámetros:
//   - id: El ID de la Persona a buscar.
//
// Retorna:
//   - Un puntero a la estructura Person.
func (r *GormPersonRepository) GetByID(id int32) (*models.Person, error) {
	var person models.Person
	err := r.db.First(&person, id).Error
	return &person, err
}

// Update actualiza los datos de una Persona en la base de datos.
// Parámetros:
//   - id: El ID de la Persona a actualizar.
//   - person: Un puntero a la estructura Person con los datos actualizados.
//
// Retorna:
//   - Un error si ocurre un problema durante la actualización.
func (r *GormPersonRepository) Update(id int32, person *models.Person) error {
	return r.db.Model(&models.Person{}).Where("id = ?", id).Updates(person).Error
}

// Delete elimina una Persona de la base de datos por su ID.
// Parámetros:
//   - id: El ID de la Persona a eliminar.
//
// Retorna:
//   - Un error si ocurre un problema durante la eliminación.
func (r *GormPersonRepository) Delete(id int32) error {
	return r.db.Delete(&models.Person{}, id).Error
}
