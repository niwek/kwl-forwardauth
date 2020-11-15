package responseform

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse ...
type UserResponse struct {
	ID        uuid.UUID       `json:"id"`
	FirstName string          `json:"firstName,omitEmpty"`
	LastName  string          `json:"lastName,omitEmpty"`
	Address   AddressResponse `json:"address,omitEmpty"`
	Dob       time.Time       `json:"dob,omitEmpty"`
	CreateAt  time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
