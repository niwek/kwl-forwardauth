package requestform

// CreateAddressRequest DTO for createAddress
type CreateAddressRequest struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitEmpty"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country,omitEmpty"`
}
