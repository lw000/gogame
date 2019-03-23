package tyrds

type Config struct {
	RedisHost     string
	RedisPassword string
}

func NewConfig() *Config {
	return &Config{}
}
