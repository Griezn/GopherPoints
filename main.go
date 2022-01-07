package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"log"
	"time"
)

//this logs in to the site and goes tot the recent points
func login() *rod.Page {
	browser := rod.New().MustConnect()
	page := browser.MustPage("https://online.myro.be/login.php")
	page.MustElement("#School").MustInput("tsaam").MustPress(input.Enter)
	_ = page.MustElement("#Username").Input("seppe.degryse@telenet.be")
	_ = page.MustElement("td:nth-child(2) > input[type=password]").Input("Sepieboy.268")
	page.MustElement("div:nth-child(3) > input[type=submit]").MustClick()
	page.MustElement("body > table > tbody > tr > td:nth-child(1) > table:nth-child(4) > tbody > tr > td > span:nth-child(2) > a").MustClick()
	log.Println("Logged in")
	return page
}

func main() {
	log.Println("start")
	page := login()
	f := readPoints(page)
	CreateVakken(f)
	seperate(f)

	time.Sleep(time.Hour)
}
