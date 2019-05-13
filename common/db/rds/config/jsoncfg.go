package ggrdsconfig

type Config struct {
	RedisHost     string
	RedisPassword string
}

func NewConfig() *Config {
	return &Config{}
}
