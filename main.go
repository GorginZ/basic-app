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

func appMetadataEndpoint(w http.ResponseWriter, r *http.Request) {

	currentCommit := getCurrentCommit()
	version := getCurrentVersionAsString()
	description := "basic app / = hello world /metadata = info"

	metadata := Metadata{Version: version, Description: description, CurrentCommit: currentCommit}
	fmt.Println("endpoint /metadata hit")
	json.NewEncoder(w).Encode(metadata)
}

func getCurrentCommit() string {
	currentCommit, ok := os.LookupEnv("CURRENT_COMMIT")
	if !ok {
		fmt.Println("CURRENT_COMMIT is not present")
	} else {
		fmt.Printf("currentCommit: %s\n", currentCommit)
	}
	return currentCommit
}

func getCurrentVersionAsString() string {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		fmt.Println("error: ", err)
	}
	versionString := string(version[:])
	return versionString
}

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func router() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/metadata", appMetadataEndpoint)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	router()
}
