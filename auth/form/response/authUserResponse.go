package responseform

import "github.com/google/uuid"

// AuthenticatedUser ...
//
// swagger:model
type AuthenticatedUser struct {
	// swagger:strfmt uuid
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Roles    []string  `json:"roles"`
}

// SwaggerAuthenticatedUser ...
//
// swagger:response authenticatedUser
type SwaggerAuthenticatedUser struct {
	// AuthenticatedUser
	//
	// in: body
	// required: true
	AuthenticatedUser AuthenticatedUser
}
