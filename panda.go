package main

import (
    "net/http"
    "log"
    "html/template"
)

// Global Variables
var hunger, health, mood int

type PageVariables struct {
    PageImage   string
    Hunger      int
    Health      int
    Mood        int
}

func main() {
    hunger = 50
    health = 50
    mood = 50
    http.HandleFunc("/", DisplayHome)
    http.HandleFunc("/feed", PandaFeed)
    http.HandleFunc("/walk", PandaWalk)
    http.HandleFunc("/play", PandaPlay)
    http.HandleFunc("/reset", PandaReset)
    log.Fatal(http.ListenAndServe(":3000", nil))
}


func DisplayHome(w http.ResponseWriter, r *http.Request) {

	// Select panda GIF
	var image string
	if hunger > 80 {
		image = "https://i.imgur.com/mxPciIj.gif"
	} else if health < 20 {
		image = "https://i.imgur.com/cbtJdl7.gif"
	} else if mood < 20 {
		image = "https://i.imgur.com/13jzL9I.gif"
	} else {
		image = "https://i.imgur.com/ojxKgtD.gif"
	}

    // Initialize display variables
    MyPageVariables := PageVariables{
        PageImage : image,
        Hunger : hunger,
        Health : health,
        Mood: mood,
    }

    // Display
    t, err := template.ParseFiles("panda.html")
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, MyPageVariables)
    if err != nil {
        log.Print("template executing error: ", err)
    }

	// Update attributes
	hunger = hunger + 5
	health = health - 5
	mood = mood - 5
	if hunger > 100 {
		hunger = 100
	}
	if health < 0 {
		health = 0
	}
	if mood < 0 {
		mood = 0
	}
}


func PandaFeed(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/s7aq5u8.gif"
    hunger = hunger - 30
    if hunger < 0 {
        hunger = 0
    }
    MyPageVariables := PageVariables{
        PageImage : image,
        Hunger : hunger,
        Health : health,
        Mood: mood,
    }

    // Display
    t, err := template.ParseFiles("panda.html")
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, MyPageVariables)
    if err != nil {
        log.Print("template executing error: ", err)
    }
}


func PandaWalk(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/f1DspuH.gif"
    health = health + 30
    if health > 100 {
        health = 100
    }
    MyPageVariables := PageVariables{
        PageImage : image,
        Hunger : hunger,
        Health : health,
        Mood: mood,
    }

    // Display
    t, err := template.ParseFiles("panda.html")
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, MyPageVariables)
    if err != nil {
        log.Print("template executing error: ", err)
    }
}


func PandaPlay(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/YgbHcz8.gif"
    mood = mood + 30
    if mood > 100 {
        mood = 100
    }
    MyPageVariables := PageVariables{
        PageImage : image,
        Hunger : hunger,
        Health : health,
        Mood: mood,
    }

    // Display
    t, err := template.ParseFiles("panda.html")
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, MyPageVariables)
    if err != nil {
        log.Print("template executing error: ", err)
    }
}


func PandaReset(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/BXzP2ym.gif"
    hunger = 50
	health = 50
	mood = 50
    MyPageVariables := PageVariables{
        PageImage : image,
        Hunger : hunger,
        Health : health,
        Mood: mood,
    }

    // Display
    t, err := template.ParseFiles("panda.html")
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, MyPageVariables)
    if err != nil {
        log.Print("template executing error: ", err)
    }
}

