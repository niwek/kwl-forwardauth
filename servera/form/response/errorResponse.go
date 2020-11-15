package responseform

// ErrorResponse return to the client whenever we have an error
//
// swagger:model
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	ModuleName string `json:"moduleName"`
}

// SwaggerErrorResponse ...
//
// swagger:response errorResponse
type SwaggerErrorResponse struct {
	// ErrorResponse
	//
	// in: body
	// required: true
	ErrorResponse ErrorResponse
}
