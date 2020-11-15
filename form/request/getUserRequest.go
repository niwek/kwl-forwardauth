package requestform

// UserID ...
//
// In route param for a user's id
// swagger:parameters getUserById
type UserID struct {
	// The users UUID as a string
	//
	// in: path
	// required:true
	UserID string `json:"userId"`
}
