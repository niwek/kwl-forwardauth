package requestform

import (
	"time"
)

// CreateUserRequest DTO
//
// swagger:model
type CreateUserRequest struct {
	FirstName string               `json:"firstName,omitEmpty"`
	LastName  string               `json:"lastName,omitEmpty"`
	Address   CreateAddressRequest `json:"address,omitEmpty"`
	Dob       time.Time            `json:"dob,omitEmpty"`
}

// SwaggerCreateUserRequest ...
//
// swagger:parameters createUser
type SwaggerCreateUserRequest struct {
	// CreateUserRequest
	//
	// in: body
	// required: true
	CreateUserRequest CreateUserRequest
}
