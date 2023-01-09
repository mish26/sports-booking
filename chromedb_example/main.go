package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	log.Println("chromedb.Run開始")

	// run task list
	var val1 string
	var val2 string
	var val3 string
	var val4 string
	var url string
	actions := []chromedp.Action{
		chromedp.Navigate(""),
		//chromedp.Navigate("http://localhost:8081/view/"),
		//chromedp.WaitVisible(`#login`, chromedp.ByID),
		//chromedp.WaitVisible("/html/body/div/form/table"),
		chromedp.WaitVisible("#childForm", chromedp.ByID),
		chromedp.WaitVisible("#login", chromedp.ByID),
		// innerHTML を取得する
		//chromedp.InnerHTML("/html/body/div/form/table", &val1),

		chromedp.Click(`#login`, chromedp.ByID),
		chromedp.WaitVisible(`/html/body`),
		// innerHTML を取得する
		chromedp.InnerHTML("/html/body", &val1),
		chromedp.WaitVisible(`#userid`, chromedp.ByID),
		chromedp.WaitVisible(`#passwd`, chromedp.ByID),
		chromedp.WaitVisible(`#login`, chromedp.ByID),
		//chromedp.SendKeys("#userid", "", chromedp.ByID),
		//chromedp.SendKeys("#passwd", "", chromedp.ByID),
		chromedp.SendKeys(`input[name="layoutChildBody:childForm:userid"]`, ""),
		chromedp.SendKeys(`input[name="layoutChildBody:childForm:passwd"]`, ""),
		chromedp.Value(`#userid`, &val3, chromedp.ByID),
		chromedp.Value(`#passwd`, &val4, chromedp.ByID),

		chromedp.Sleep(3 * time.Second),
		//chromedp.Click(`#login`, chromedp.ByID),
		//chromedp.Click(`#login`, chromedp.ByID),
		//chromedp.Submit(`//input[@value="ログイン"]`),
		chromedp.Submit("#login", chromedp.ByID),
		chromedp.Sleep(4 * time.Second),
		//chromedp.WaitVisible(`/html/body`),
		chromedp.InnerHTML("/html/body", &val2),
		// innerHTML を取得する
		//chromedp.InnerHTML("/html/body", &val2),

		chromedp.Location(&url),
		//chromedp.Value(`#loginJKey`, &val1, chromedp.ByID),
	}
	if err := chromedp.Run(ctx, actions...); err != nil {
		log.Fatalln(err)
	}
	log.Println("val", val1)
	log.Println("========================")
	log.Println("========================")
	log.Println("val2", val2)
	log.Println("========================")
	log.Println("========================")
	log.Println("val3, val4", val3, val4)
	log.Println("========================")
	log.Println("========================")
	log.Println(url)
	log.Println("chromedb.Run終了")

}
