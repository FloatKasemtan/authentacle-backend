package config

type config struct {
	JwtSecret     string   `yaml:"JWT_SECRET"`
	Cors          []string `yaml:"CORS"`
	Port          string   `yaml:"PORT"`
	DbHost        string   `yaml:"DB_HOST"`
	DbName        string   `yaml:"DB_NAME"`
	EncryptionKey string   `yaml:"ENCRYPTION_KEY"`
	RedisHost     string   `yaml:"REDIS_HOST"`
	RedisPassword string   `yaml:"REDIS_PASSWORD"`
	RedisDb       int      `yaml:"REDIS_DB"`
}
