package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/joho/godotenv"
	"sports-booking/cmd/config"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("%+v\n", err)
	}

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	// Headless runs the browser on foreground, you can also use flag "-rod=show"
	// Devtools opens the tab in each new tab opened automatically
	//l := launcher.New().
	//	Headless(false).
	//	Devtools(true)
	////
	//defer l.Cleanup() // remove launcher.FlagUserDataDir
	////
	//url := l.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// SlowMotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
	browser := rod.New().
		//ControlURL(url).
		//Trace(true).
		SlowMotion(2 * time.Second).
		MustConnect()

	defer browser.MustClose()

	page := browser.MustPage(cfg.URL)

	page.MustElement("#login").MustType(input.Enter)

	page.MustElement("#userid").MustInput(cfg.UserID)
	page.MustElement("#passwd").MustInput(cfg.Password).MustType(input.Enter)
	text := page.MustElement("#childForm").MustText()

	fmt.Println(text)
}
