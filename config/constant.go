package config

type constant struct {
	Env            string
	AllowedOrigins string
}

// Constant ...
var Constant = constant{
	Env:            "ENV",
	AllowedOrigins: "ALLOWED_ORIGINS",
}
