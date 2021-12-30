package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"time"
)


//this logs in to the site and goes tot the recent points
func login(page *rod.Page){
	page.MustElement("#School").MustInput("tsaam").MustPress(input.Enter)
	page.MustElement("#SmsRefB").MustClick()
	page.MustElement("#login_form__username").MustWaitLoad().MustInput("seppe.degryse@student.tsaam.be")
	page.MustElement("#login_form__password").MustInput("Sepieboy.123")
	page.MustElement("#smscMain > div > div > div.oauth__form > form > button").MustClick()
	page.MustElement("body > table > tbody > tr > td:nth-child(1) > table > tbody > tr > td > span:nth-child(1) > a").MustClick()//clicks on logbook
	page.MustElement("body > table > tbody > tr > td:nth-child(1) > table:nth-child(4) > tbody > tr > td > span:nth-child(2) > a").MustClick()//clicks on recent
}

func main() {
	browser := rod.New().MustConnect()
	page := browser.MustPage("https://online.myro.be/login.php")
	login(page)


	time.Sleep(time.Hour)
}