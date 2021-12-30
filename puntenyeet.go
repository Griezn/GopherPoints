package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-rod/rod"
)

type points struct {
	vak       string
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
	//page.MustElement("#OptionTable > tbody > tr > td.countInput > input").MustSelectAllText().MustInput("").MustInput("50")

	for i := 1; i < 20; i++ {
		//search for the names and add them to the struct
		course := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.courseNameRecent", i)).MustText()
		subject := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.title", i)).MustText()
		punten := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.point", i)).MustText()
		puntenmax := page.MustElement(fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.max", i)).MustText()

		puntenlijst = append(puntenlijst, points{
			vak:       course,
			onderwerp: subject,
			points:    punten,
			pointsmax: puntenmax,
		})
	}
	return puntenlijst
}
