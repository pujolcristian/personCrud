package repositories

import "personCrud/internal/domain/models"

type PersonRepository interface {
	Create(person *models.Person) error
	GetAll() ([]models.Person, error)
	GetByID(id int32) (*models.Person, error)
	Update(id int32, person *models.Person) error
	Delete(id int32) error
}
