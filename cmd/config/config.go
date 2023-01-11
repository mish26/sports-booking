package config

type Config struct {
	URL      string `env:"URL"`
	UserID   int    `env:"USER_ID"`
	Password int    `env:"PASSWORD"`
}
