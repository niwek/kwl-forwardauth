package config

type constant struct {
	Env            string
	AllowedOrigins string
	Port           string
}

// Constant ...
var Constant = constant{
	Env:            "ENV",
	AllowedOrigins: "ALLOWED_ORIGINS",
	Port:           "PORT",
}
