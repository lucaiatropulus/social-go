package config

type dbConfig struct {
	Address            string `json:"address"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
	MaxIdleConnections int    `json:"maxIdleConnections"`
	MaxIdleTime        string `json:"maxIdleTime"`
}
