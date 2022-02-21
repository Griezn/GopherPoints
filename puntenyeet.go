package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	_ "github.com/go-sql-driver/mysql"
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

type user struct {
	user_id  int
	email    string
	password string
}

//this logs in to the site and goes tot the recent points
func (u *user) login() *rod.Page {
	browser := rod.New().MustConnect()
	page := browser.MustPage("https://online.myro.be/login.php")
	page.MustElement("#School").MustInput("tsaam").MustPress(input.Enter)
	_ = page.MustElement("#Username").Input(u.email)
	_ = page.MustElement("td:nth-child(2) > input[type=password]").Input(u.password)
	page.MustElement("div:nth-child(3) > input[type=submit]").MustClick()
	page.MustElement("body > table > tbody > tr > td:nth-child(1) > table:nth-child(4) > tbody > tr > td > span:nth-child(2) > a").MustClick()
	log.Printf("Logged in as %v", u.email)
	return page
}

//read all the latest 50 points handy for first time login
func (u *user) readPoints() []points {
	page := u.login()
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

		new_subject := strings.ReplaceAll(subject, "'", "''")

		puntenlijst = append(puntenlijst, points{
			Vak: course,
			Test: toets{
				Onderwerp: new_subject,
				Points:    punten,
				Pointsmax: puntenmax,
				Procent:   *procent,
			},
		})
	}
	log.Println("End search")

	return puntenlijst
}

//get only the latest result
func (u *user) readLatest() points {
	page := u.login()
	//array to return
	var latestPoint points
	page.MustElement(`[type="text"]`).MustSelectAllText().MustPress(input.Backspace).MustInput("50")

	log.Println("Start search")
	//search for the names and add them to the struct points
	course := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(1) > td.courseNameRecent").MustText()
	subject := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(1) > td.title").MustText()
	punten := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(1) > td.point").MustText()
	puntenmax := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(1) > td.max").MustText()
	procent := page.MustElement("#\\31 7_2530 > tbody > tr:nth-child(1) > td.point").MustAttribute("title")

	latestPoint = points{
		Vak: course,
		Test: toets{
			Onderwerp: subject,
			Points:    punten,
			Pointsmax: puntenmax,
			Procent:   *procent,
		},
	}
	log.Println("End search")
	log.Println(latestPoint)

	return latestPoint
}

//save the list of points as JSON if we want to
func (u *user) SaveJSON() {
	lijst := u.readPoints()

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

//Read the points from myro and add them to the database
func (u *user) updateDatabse(db *sql.DB) {
	list := u.readPoints()

	log.Println("Start insert")
	for _, v := range list {
		insert, err := db.Query(fmt.Sprintf("INSERT INTO points (owner_id, vak, onderwerp, points, pointsmax, procent) VALUES (%d, '%v', '%v', '%v', '%v', '%v');",
			u.user_id, v.Vak, v.Test.Onderwerp, v.Test.Points, v.Test.Pointsmax, v.Test.Procent))
		if err != nil {
			log.Fatalf("Could not insert! Error; %v", err)
		}
		defer insert.Close()
	}
	log.Println("End insert")
}

func (u *user) firstlogin() bool {
	return true
}

func main() {
	//connect to database
	db, err := sql.Open("mysql", "test:password@tcp(localhost:3306)/gopherpoints")
	if err != nil {
		log.Println("Could not open SQL")
	}
	defer db.Close()

	//get the users and add them to a list
	var userList []user
	users, err := db.Query("SELECT id, email, ppword FROM users;")
	if err != nil {
		log.Fatalf("Could not get the users")
	}
	for users.Next() {
		var id int
		var email, password string
		users.Scan(&id, &email, &password)

		userList = append(userList, user{
			user_id:  id,
			email:    email,
			password: password,
		})
	}

	//every 20 seconds (for now) loop over the userList and update the database
	for range time.Tick(time.Second * 20) {
		log.Println("yeet")
		for _, u := range userList {
			if u.firstlogin() {
				u.updateDatabse(db)
			} else {
				u.readLatest()
			}
		}
	}
}
