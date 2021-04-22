package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Event struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/newevent", createNewEvent)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewEvent(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var event Event
	json.Unmarshal(reqBody, &event)
	// update our global Articles array to include
	// our new Article
	fmt.Println(event)

}
