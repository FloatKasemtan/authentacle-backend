package config

type config struct {
	JWT_SECRET string   `yaml:"JWT_SECRET"`
	CORS       []string `yaml:"CORS"`
	PORT       string   `yaml:"PORT"`
	DB_HOST    string   `yaml:"DB_HOST"`
}
