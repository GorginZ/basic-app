package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	currentCommit, ok := os.LookupEnv("CURRENT_COMMIT")
	if !ok {
		fmt.Println("CURRENT_COMMIT is not present")
	} else {
		fmt.Printf("currentCommit: %s\n", currentCommit)
	}

	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		fmt.Println("error: ", err)
	}
	//figure out error handling here
	// if err == nil {
	versionString := string(version[:])
	// }

	metadata := Metadata{Version: versionString, Description: "basic app", CurrentCommit: currentCommit}
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
