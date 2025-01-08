package usecases

import (
	"fmt"
	"personCrud/internal/domain/dtos"
	"personCrud/internal/domain/mappers"
	"personCrud/internal/domain/models"
	"personCrud/internal/domain/repositories"
	"personCrud/internal/domain/validators"
)

type PersonUseCase struct {
	repo      repositories.PersonRepository
	validator *validators.PersonValidator
}

func NewPersonUseCase(repo repositories.PersonRepository) *PersonUseCase {
	return &PersonUseCase{
		repo:      repo,
		validator: validators.NewPersonValidator(),
	}
}

// CreatePerson crea una nueva Persona en el sistema.
// Realiza las validaciones necesarias antes de enviar los datos al repositorio.
// Parámetros:
//   - person: Un puntero a la estructura Person que contiene los datos de la persona.
//
// Retorna:
//   - Un error si ocurre un problema durante la validación o la creación.
func (uc *PersonUseCase) CreatePerson(person *models.Person) error {

	// Validar el modelo antes de procesarlo
	if err := uc.validator.Validate(person); err != nil {
		return err
	}

	// Verificar si el usuario ya existe
	_, err := uc.repo.GetByID(person.ID)
	if err == nil {
		return fmt.Errorf("user with ID %d already exists", person.ID)
	}

	// Guardar en la base de datos
	return uc.repo.Create(person)
}

// GetAllPersons obtiene una lista de todas las Personas registradas en el sistema.
// Retorna:
//   - Una lista de Personas en formato DTO.
func (uc *PersonUseCase) GetAllPersons() ([]dtos.PersonDTO, error) {

	persons, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return mappers.MapPersonsToDTO(persons), nil
}

// GetPersonByID obtiene una Persona por su ID.
// Parámetros:
//   - id: El ID de la Persona a buscar.
//
// Retorna:
//   - Un puntero a la estructura Person en formato DTO.
func (uc *PersonUseCase) GetPersonByID(id int32) (*dtos.PersonDTO, error) {
	person, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return mappers.MapPersonToDTO(person), nil
}

func (uc *PersonUseCase) UpdatePerson(id int32, person *models.Person) error {

	if err := uc.validator.Validate(person); err != nil {
		return err
	}

	return uc.repo.Update(id, person)
}

func (uc *PersonUseCase) DeletePerson(id int32) error {
	return uc.repo.Delete(id)
}
