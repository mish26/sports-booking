package config

type Config struct {
	URL      string `env:"URL"`
	UserID   string `env:"USER_ID"`
	Password string `env:"PASSWORD"`
}
