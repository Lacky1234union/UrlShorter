package config

type Config struct {
	Env string `yaml:"env" env:"ENV" env-default:"local"`
}
