package config

type Config struct {
	ServerConfig *ServerConfig `json:"server"`
	Database     *Database     `json:"database"`
	AppConfig    *AppConfig    `json:"appConfig"`
}

type ServerConfig struct {
}

type Database struct {
}

type AppConfig struct {
}

func NewConfigServer() *Config {
	return &Config{
		ServerConfig: &ServerConfig{},
		Database:     &Database{},
		AppConfig:    &AppConfig{},
	}
}
