package main

import (
	"fmt"
	"github.com/go-rod/rod"
)

type points struct {
	vak string
	onderwerp string
	points string
	pointsmax string
}

func readPoints(page *rod.Page) []points{
	var puntenlijst []points

	//loop parameters
	end := false
	i := 1

	course := fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.courseNameRecent", i)
	subject := fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.title", i)
	punten := fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.point", i)
	puntenmax := fmt.Sprintf("#\\31 7_2530 > tbody > tr:nth-child(%d) > td.max", i)

	for !end{

	}


	//vak := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(5) > td.courseNameRecent").MustText()
	//toets := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(5) > td.title").MustText()
	//punten := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(5) > td.point").MustText()
	//puntenmax := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(5) > td.max").MustText()
	return puntenlijst
}

/* compare title
#\31 7_2530 > tbody > tr:nth-child(1) > td.courseNameRecent
#\31 7_2530 > tbody > tr:nth-child(5) > td.courseNameRecent
#\31 7_2530 > tbody > tr:nth-child(6) > td.courseNameRecent

compare points
#\31 7_2530 > tbody > tr:nth-child(12) > td.point  //genoeg
#\31 7_2530 > tbody > tr:nth-child(13) > td.point.tekort  //tekort

document.querySelector("#\\31 7_2530 > tbody > tr:nth-child(12) > td.point")
document.querySelector("#\\31 7_2530 > tbody > tr:nth-child(13) > td.point.tekort")

 */