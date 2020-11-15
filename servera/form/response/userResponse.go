package responseform

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse ...
//
// swagger:model
type UserResponse struct {
	// swagger:strfmt uuid
	ID        uuid.UUID       `json:"id"`
	FirstName string          `json:"firstName,omitEmpty"`
	LastName  string          `json:"lastName,omitEmpty"`
	Address   AddressResponse `json:"address,omitEmpty"`
	Dob       time.Time       `json:"dob,omitEmpty"`
	CreateAt  time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

// SwaggerUserResponse ...
//
// swagger:response userResponse
type SwaggerUserResponse struct {
	// UserResponse
	//
	// in: body
	// required: true
	UserResponse UserResponse
}
