package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/joho/godotenv"
	"sports-booking/cmd/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("%+v\n", err)
	}

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	browser := rod.New().MustConnect()

	defer browser.MustClose()

	page := browser.MustPage(cfg.URL)

	page.MustElement("#login").MustType(input.Enter)

	page.MustElement("#userid").MustInput(string(cfg.UserID))
	page.MustElement("#userid").MustInput(string(cfg.Password)).MustType(input.Enter)

	text := page.MustElement("#childForm").MustText()

	fmt.Println(text)
}
