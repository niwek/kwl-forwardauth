package responseform

// AddressResponse DTO for createAddress
//
// swagger:model
type AddressResponse struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitEmpty"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country,omitEmpty"`
}

// SwaggerAddressResponse ...
//
// swagger:response addressResponse
type SwaggerAddressResponse struct {
	// AddressResponse
	//
	// in: body
	// required: true
	AddressResponse AddressResponse
}
