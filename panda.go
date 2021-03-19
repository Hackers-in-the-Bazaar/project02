package main

import (
    "net/http"
    "log"
    "html/template"
	"time"
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
    //http.HandleFunc("/selected", UserSelected)
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func DisplayHome(w http.ResponseWriter, r *http.Request) {

	// Select panda GIF
	var image string
	if hunger > 80 {
		image = "https://i.imgur.com/UtVAPiM.jpeg"		// TODO: change later
	} else if health < 20 {
		image = "https://i.imgur.com/UtVAPiM.jpeg"		// TODO: change later
	} else if mood < 20 {
		image = "https://i.imgur.com/UtVAPiM.jpeg"		// TODO: change later
	} else {
		image = "https://i.imgur.com/BXzP2ym.gif"
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

	// Update attributes and refresh
	//time.Sleep(20 * time.Second)
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
	http.Redirect(w, r, "http://3.141.45.193:3000", 301)
	//http.Redirect(w, r, "/", 307)
}


func PandaFeed(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/s7aq5u8.gif"
    hunger = hunger - 20
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

	// Redirect to homepage
	time.Sleep(6 * time.Second)
	http.Redirect(w, r, "http://3.141.45.193:3000", 307)
}


func PandaWalk(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/f1DspuH.gif"
    health = health + 20
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

	// Redirect to homepage
	time.Sleep(6 * time.Second)
	http.Redirect(w, r, "http://3.141.45.193:3000", 307)
}


func PandaPlay(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/YgbHcz8.gif"
    mood = mood + 20
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

	// Redirect to homepage
	time.Sleep(6 * time.Second)
	http.Redirect(w, r, "http://3.141.45.193:3000", 307)
}

