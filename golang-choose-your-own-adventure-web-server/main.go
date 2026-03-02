package main

import (
	"choose-your-own-adventure/cyoa"
	"log"
	"net/http"
	"os"
)

func main() {

	//read the json file
	filename := "cyoa.json"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", story)

	//set up the router
	h := cyoa.NewHandler(story)
	http.ListenAndServe("localhost:3000", h)

}
