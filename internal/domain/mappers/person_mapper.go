package mappers

import (
	"personCrud/internal/domain/dtos"
	"personCrud/internal/domain/models"
)

func MapPersonToDTO(person *models.Person) *dtos.PersonDTO {
	return &dtos.PersonDTO{
		ID:      person.ID,
		Name:    person.Name,
		Address: person.Address,
		Phone:   person.Phone,
		Email:   person.Email,
	}
}

func MapPersonsToDTO(persons []models.Person) []dtos.PersonDTO {
	dto := make([]dtos.PersonDTO, len(persons))
	for i, person := range persons {
		dto[i] = *MapPersonToDTO(&person)
	}
	return dto
}
