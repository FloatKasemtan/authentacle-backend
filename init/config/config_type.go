package config

type config struct {
	JWT_SECRET     string   `yaml:"JWT_SECRET"`
	CORS           []string `yaml:"CORS"`
	PORT           string   `yaml:"PORT"`
	DB_HOST        string   `yaml:"DB_HOST"`
	DB_NAME        string   `yaml:"DB_NAME"`
	ENCRYPTION_KEY string   `yaml:"ENCRYPTION_KEY"`
}
