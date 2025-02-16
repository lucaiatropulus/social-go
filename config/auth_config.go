package config

type authConfig struct {
	Secret        string `json:"secret"`
	ValidDuration string `json:"validDuration"`
	Issuer        string `json:"issuer"`
	Audience      string `json:"audience"`
}
