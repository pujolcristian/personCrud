package models

import (
	"time"
)

// Person representa una entidad de Persona en el sistema.
// Contiene los datos personales que se guardan en la base de datos.
type Person struct {
	ID        int32     `gorm:"primaryKey;column:id" json:"id" validate:"required,numeric"`
	Name      string    `gorm:"column:name" json:"name" validate:"required,min=3,max=50"`                            // String with length constraints
	Address   string    `gorm:"column:address" json:"address" validate:"required,min=5,max=100"`                     // String with length constraints
	Phone     uint      `gorm:"column:phone" json:"phone" validate:"required,numeric,min=1000000000,max=9999999999"` // Numeric with exact digits
	Email     string    `gorm:"column:email" json:"email" validate:"omitempty,email"`                                // Valid email format
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`                                    // Must be a valid datetime
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime:milli" json:"updatedAt"`                              // Must be a valid datetime
}

func (Person) TableName() string {
	return "persons"
}
