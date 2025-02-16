package config

type appConfig struct {
	Address     string `json:"address"`
	Environment string `json:"environment"`
	ApiURL      string `json:"apiURL"`
	FrontendURL string `json:"frontendURL"`
}
