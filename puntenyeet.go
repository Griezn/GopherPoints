package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

type points struct {
	Vak  string `json:"vak"`
	Test toets  `json:"toets"`
}

type toets struct {
	Onderwerp string `json:"onderwerp"`
	Points    string `json:"points"`
	Pointsmax string `json:"pointsmax"`
	Procent   string `json:"procent"`
}

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

func readPoints(page *rod.Page) []points {
	//array to return
	var puntenlijst []points
	page.MustElement(`[type="text"]`).MustSelectAllText().MustPress(input.Backspace).MustInput("50")

	log.Println("Start search")
	for i := 1; i < 50; i++ {
		//search for the names and add them to the struct points
		course := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.courseNameRecent", i)).MustText()
		subject := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.title", i)).MustText()
		punten := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.point", i)).MustText()
		puntenmax := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.max", i)).MustText()
		procent := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.point", i)).MustAttribute("title")

		puntenlijst = append(puntenlijst, points{
			Vak: course,
			Test: toets{
				Onderwerp: subject,
				Points:    punten,
				Pointsmax: puntenmax,
				Procent:   *procent,
			},
		})
	}
	log.Println("End search")

	//Save the points as a JSON file

	return puntenlijst
}

func SaveJSON() {
	page := login()
	lijst := readPoints(page)

	//write json
	data, err := json.MarshalIndent(&lijst, "", " ")
	if err != nil {
		log.Println(err)
	}

	//write to file
	err = ioutil.WriteFile("./server/yeet.json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	SaveJSON()
}
