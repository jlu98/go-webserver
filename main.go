package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// this is an important comment
// another important comment
type posts struct {
	UserId int    `json:"userId`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Welcome to my exercise server!\n")
}
func getPing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /ping request\n")
	io.WriteString(w, "Pong!\n")
}
func getAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /API request\n")
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	textBytes := []byte(responseData)
	post1 := posts{}

	err = json.Unmarshal(textBytes, &post1)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.WriteString(w, fmt.Sprintf("UserId: %d\nPostId: %d\nTitle:\n%s\nBody:\n%s\n", post1.UserId, post1.Id, post1.Title, post1.Body))
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/ping", getPing)
	http.HandleFunc("/api", getAPI)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
