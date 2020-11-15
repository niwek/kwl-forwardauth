package responseform

// ErrorResponse return to the client whenever we have an error
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	ModuleName string `json:"moduleName"`
}
