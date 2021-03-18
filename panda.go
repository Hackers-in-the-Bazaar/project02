package main

import (
	"fmt"
    "net/http"
    "log"
    "html/template"
    "os"
)

type RadioButton struct {
    Name       string
    Value      string
    IsDisabled bool
    IsChecked  bool
    Text       string
}

type PageVariables struct {
  PageImageName    string
  PageRadioButtons []RadioButton
  Answer           string
}

func main() {
  http.HandleFunc("/", DisplayHome)
  http.HandleFunc("/selected", UserSelected)
  log.Fatal(http.ListenAndServe(":3000", nil))
}

func DisplayHome(w http.ResponseWriter, r *http.Request) {

    // Display buttons
	ImageName := "images/tenor.gif"
    _, err := os.Stat(ImageName)
	MyRadioButtons := []RadioButton{
    RadioButton{"actionselect", "cats", false, false, "Cats"},
    RadioButton{"actionselect", "dogs", false, false, "Dogs"},
    }


  MyPageVariables := PageVariables{
    PageImageName : ImageName,
    PageRadioButtons : MyRadioButtons,
    }

   t, err := template.ParseFiles("panda.html") //parse the html file homepage.html
   if err != nil { // if there is an error
     log.Print("template parsing error: ", err) // log it
   }

   err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
   if err != nil { // if there is an error
     log.Print("template executing error: ", err) //log it
   }

}

func UserSelected(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  // r.Form is now either
  // map[animalselect:[cats]] OR
  // map[animalselect:[dogs]]
 // so get the animal which has been selected
  youranimal := r.Form.Get("animalselect")

  MyPageVariables := PageVariables{
    Answer : youranimal,
    }

 // generate page by passing page variables into template
    t, err := template.ParseFiles("panda.html") //parse the html file homepage.html
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
}
