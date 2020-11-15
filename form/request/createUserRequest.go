package requestform

import (
	"time"
)

// CreateUserRequest DTO
type CreateUserRequest struct {
	FirstName string               `json:"firstName,omitEmpty"`
	LastName  string               `json:"lastName,omitEmpty"`
	Address   CreateAddressRequest `json:"address,omitEmpty"`
	Dob       time.Time            `json:"dob,omitEmpty"`
}
