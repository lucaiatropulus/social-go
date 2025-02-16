package config

type mailConfig struct {
	ApiKey     string `json:"apiKey"`
	Email      string `json:"email"`
	Expiration string `json:"expiration"`
}
