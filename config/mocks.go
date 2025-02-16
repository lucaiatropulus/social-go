package config

func NewMockConfig() *Config {
	return &Config{
		APP: appConfig{Environment: "development"},
		Auth: authConfig{
			Secret:        "testing_secret",
			ValidDuration: "5m",
			Issuer:        "testing_issuer",
			Audience:      "testing_audience",
		},
	}
}
