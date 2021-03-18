package main 

import (
  "fmt"
	"net/http"
	"os"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
  http.HandleFunc("/", test)
	http.ListenAndServe(getPort(), nil)
}

func test(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "this is a test")
}
