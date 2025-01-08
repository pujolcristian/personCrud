package dtos

type PersonDTO struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   uint   `json:"phone"`
	Email   string `json:"email"`
}
