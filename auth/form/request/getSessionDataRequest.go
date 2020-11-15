package requestform

// Token ...
//
// In route param for a user's id
// swagger:parameters getSessionData
type Token struct {
	// Token
	//
	// in: path
	// required:true
	Token string `json:"token"`
}

// Authorization ...
//
// Authorization Bearer token
// swagger:parameters getSessionDataFromHeader
type Authorization struct {
	// Token
	//
	// in: header
	// required:true
	Authorization string `json:"authorization"`
}
