package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	URL      string `env:"URL"`
	UserID   int    `env:"USER_ID"`
	Password int    `env:"PASSWORD"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("%+v\n", err)
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
}
