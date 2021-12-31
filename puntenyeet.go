package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"log"
)

type points struct {
	vak  string
	test toets
}

type toets struct {
	onderwerp string
	points    string
	pointsmax string
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
			vak: course,
			test: toets{
				onderwerp: subject,
				points:    punten,
				pointsmax: puntenmax,
			},
		})
	}
	log.Println("End search")
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
		vakken = append(vakken, point.vak)
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
		switch point.vak {
		case "Engels":
			Engels = append(Engels, point.test)
			i++
			continue
		case "Aardrijkskunde":
			Aardrijkskunde = append(Aardrijkskunde, point.test)
			i++
			continue
		case "Fysica":
			Fysica = append(Fysica, point.test)
			i++
			continue
		case "Frans":
			Frans = append(Frans, point.test)
			i++
			continue
		case "Duits":
			Duits = append(Duits, point.test)
			i++
			continue
		case "Nederlands":
			Nederlands = append(Nederlands, point.test)
			i++
			continue
		case "Biologie":
			Biologie = append(Biologie, point.test)
			i++
			continue
		case "Chemie":
			Chemie = append(Chemie, point.test)
			i++
			continue
		case "Esthetica":
			Esthetica = append(Esthetica, point.test)
			i++
			continue
		case "Godsdienst":
			Godsdienst = append(Godsdienst, point.test)
			i++
			continue
		case "Wiskunde":
			Wiskunde = append(Wiskunde, point.test)
			i++
			continue
		default:
			restjes = append(restjes, point.test)
			i++
			continue
		}
	}
	fmt.Println(Wiskunde)
}
