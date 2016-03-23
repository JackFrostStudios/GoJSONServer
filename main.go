package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type serverResponse struct {
	TimeDelay int32
}

var maxDelay int32

func main() {
	fmt.Printf("Please enter the maximum delay to be tested: ")
	fmt.Scan(&maxDelay)
	fmt.Printf("Listening on localhost:9001, will respond after random delay with JSON.\n")
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", returnJSON)         // set router
	err := http.ListenAndServe(":9001", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func returnJSON(w http.ResponseWriter, r *http.Request) {
	var timeDelay int32
	if maxDelay == 0 {
		timeDelay = 0
	} else {
		timeDelay = rand.Int31n(maxDelay)
	}

	response := serverResponse{timeDelay}
	fmt.Printf("Sleeping %d ....\n", timeDelay)
	time.Sleep(time.Second * time.Duration(timeDelay))

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Returning %d ....\n", timeDelay)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
