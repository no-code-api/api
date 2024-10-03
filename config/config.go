package config

type Config struct {
	ServerPort      string
	PostgreHost     string
	PostgrePort     string
	PostgreUserName string
	PostgrePassword string
	PostgreDbName   string
	PostgreSSLMode  string
}

var Env *Config

func Initialize() {
	Env = LoadEnv()
}
