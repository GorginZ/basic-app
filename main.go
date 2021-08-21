package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Metadata struct {
	Version       string `json:"Version"`
	Description   string `json:"Description"`
	CurrentCommit string `json:"CurrentCommit"`
}

func appMetadata(w http.ResponseWriter, r *http.Request) {
	// currentCommit := os.Getenv("CURRENT_COMMIT")

	currentCommit, ok := os.LookupEnv("CURRENT_COMMIT")
	if !ok {
		fmt.Println("CURRENT_COMMIT is not present")
	} else {
		fmt.Printf("currentCommit: %s\n", currentCommit)
	}

	metadata := Metadata{Version: "1.0.0", Description: "basic app", CurrentCommit: currentCommit}
	fmt.Println("endpoint /metadata hit")
	json.NewEncoder(w).Encode(metadata)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/metadata", appMetadata)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	handleRequests()
}
