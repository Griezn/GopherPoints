package main

import (
	"context"
	"encoding/json"
	"errors"
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
}

func handleError(err error) {
	var evalErr *rod.ErrEval
	if errors.Is(err, context.DeadlineExceeded) { // timeout error
		fmt.Println("timeout err")
	} else if errors.As(err, &evalErr) { // eval error
		fmt.Println(evalErr.LineNumber)
	} else if err != nil {
		fmt.Println("can't handle", err)
	}
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

		puntenlijst = append(puntenlijst, points{
			Vak: course,
			Test: toets{
				Onderwerp: subject,
				Points:    punten,
				Pointsmax: puntenmax,
			},
		})
	}
	log.Println("End search")

	data, err := json.MarshalIndent(puntenlijst, "", " ")
	if err != nil {
		log.Println(err)
	}
	err2 := ioutil.WriteFile("yeet.json", data, 0644)
	if err2 != nil {
		log.Println(err2)
	}

	return puntenlijst
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func CreateVakken(points []points) {
	var vakken []string
	for i, point := range points {
		vakken = append(vakken, point.Vak)
		i++
	}
	vakken = removeDuplicateStr(vakken)
	fmt.Println(vakken)
}

func seperate(points []points) {
	var Engels []toets
	var Aardrijkskunde []toets
	var Fysica []toets
	var Frans []toets
	var Duits []toets
	var Nederlands []toets
	var Biologie []toets
	var Chemie []toets
	var Esthetica []toets
	var Godsdienst []toets
	var Wiskunde []toets
	var restjes []toets

	for i, point := range points {
		switch point.Vak {
		case "Engels":
			Engels = append(Engels, point.Test)
			i++
			continue
		case "Aardrijkskunde":
			Aardrijkskunde = append(Aardrijkskunde, point.Test)
			i++
			continue
		case "Fysica":
			Fysica = append(Fysica, point.Test)
			i++
			continue
		case "Frans":
			Frans = append(Frans, point.Test)
			i++
			continue
		case "Duits":
			Duits = append(Duits, point.Test)
			i++
			continue
		case "Nederlands":
			Nederlands = append(Nederlands, point.Test)
			i++
			continue
		case "Biologie":
			Biologie = append(Biologie, point.Test)
			i++
			continue
		case "Chemie":
			Chemie = append(Chemie, point.Test)
			i++
			continue
		case "Esthetica":
			Esthetica = append(Esthetica, point.Test)
			i++
			continue
		case "Godsdienst":
			Godsdienst = append(Godsdienst, point.Test)
			i++
			continue
		case "Wiskunde":
			Wiskunde = append(Wiskunde, point.Test)
			i++
			continue
		default:
			restjes = append(restjes, point.Test)
			i++
			continue
		}
	}
	//fmt.Println(Wiskunde)
}
