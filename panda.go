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
    //http.HandleFunc("/selected", UserSelected)
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func DisplayHome(w http.ResponseWriter, r *http.Request) {

    // Initialize display variables
	image := "https://i.imgur.com/UtVAPiM.jpeg"
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
}

