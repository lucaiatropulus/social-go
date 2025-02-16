package config

type redisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Database int    `json:"database"`
	Enabled  bool   `json:"enabled"`
}
